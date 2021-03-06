
const TILE_SIZE = 30;
const FPS = 30;
const SLEEP = 1000 / FPS;

enum RawTile {
  AIR,
  FLUX,
  UNBREAKABLE,
  PLAYER,
  STONE, FALLING_STONE,
  BOX, FALLING_BOX,
  KEY1, LOCK1,
  KEY2, LOCK2
}

enum RawInput {
  UP, DOWN, LEFT, RIGHT
}

let playerx = 1;
let playery = 1;

let rawMap: RawTile[][] = [
  [2, 2, 2, 2, 2, 2, 2, 2],
  [2, 3, 0, 1, 1, 2, 0, 2],
  [2, 4, 2, 6, 1, 2, 0, 2],
  [2, 8, 4, 1, 1, 2, 0, 2],
  [2, 4, 1, 1, 1, 9, 0, 2],
  [2, 2, 2, 2, 2, 2, 2, 2],
];

let map: Tile[][];
function assertExhausted(x: never): never {
  throw new Error("Unexpected object: " + x);
}

function transformTile(tile: RawTile) {
  switch (tile) {
    case RawTile.AIR: return new Air();
    case RawTile.PLAYER: return new Player();
    case RawTile.UNBREAKABLE: return new Unbreakable();
    case RawTile.STONE: return new Stone(false);
    case RawTile.FALLING_STONE: return new Stone(true);
    case RawTile.BOX: return new Box(false);
    case RawTile.FALLING_BOX: return new Box(true);
    case RawTile.FLUX: return new Flux();
    case RawTile.KEY1: return new Key1();
    case RawTile.LOCK1: return new Lock1();
    case RawTile.KEY2: return new Key2();
    case RawTile.LOCK2: return new Lock2();
    default: assertExhausted(tile);
  }
}

function transformMap() {
  map = new Array(rawMap.length);
  for(let y = 0; y < rawMap.length; y++) {
    map[y] = new Array(rawMap[y].length);
    for(let x = 0; x < rawMap[y].length; x++) {
      map[y][x] = transformTile(rawMap[y][x]);
    }
  }
}

let inputs: Input[] = [];

function removeLock1() {
  for (let y = 0; y < map.length; y++) {
    for (let x = 0; x < map[y].length; x++) {
      if (map[y][x].isLock1()) {
        map[y][x] = new Air();
      }
    }
  }
}

function removeLock2() {
  for (let y = 0; y < map.length; y++) {
    for (let x = 0; x < map[y].length; x++) {
      if (map[y][x].isLock2()) {
        map[y][x] = new Air();
      }
    }
  }
}

function moveToTile(newx: number, newy: number) {
  map[playery][playerx] = new Air(); 
  map[newy][newx] = new Player(); //Tile.PLAYER;
  playerx = newx;
  playery = newy;
}

function moveHorizontal(dx: number) {
  map[playery][playerx + dx].moveHorizontal(dx);
}

function moveVertical(dy: number) {
  map[playery + dy][playerx].moveVertical(dy);
}

function update() {
  handleInputs();
  updateMap();
}

function handleInputs() {
  while (inputs.length > 0) {
    let input = inputs.pop();
    input.handle();
  }
}

interface Input {
  handle(): void;
}

class Right implements Input {
  handle() { moveHorizontal(1); }
}

class Left implements Input {
  handle() { moveHorizontal(-1); }
}

class Up implements Input {
  handle() { moveVertical(-1); }
}

class Down implements Input {
  handle() { moveVertical(1); }
}

function updateMap() {
  for (let y = map.length - 1; y >= 0; y--) {
    for (let x = 0; x < map[y].length; x++) {
      updateTile( x, y );
    }
  }
}

function updateTile( x: number, y: number ) {
  map[y][x].update(x, y);
}

function createGraphics() {
  let canvas = document.getElementById("GameCanvas") as HTMLCanvasElement;
  let g = canvas.getContext("2d");
  g.clearRect(0, 0, canvas.width, canvas.height);
  return g;
}

function draw() {
  let g = createGraphics();
  drawMap( g );
  drawPlayer( g );
}

function drawMap( g: CanvasRenderingContext2D ) {
  for (let y = 0; y < map.length; y++) {
    for (let x = 0; x < map[y].length; x++) {
      map[y][x].draw( g, x, y );
    }
  }
}

interface Tile {
  isAir(): boolean;
  isLock1(): boolean;
  isLock2(): boolean;

  moveHorizontal(dx: number): void;
  moveVertical(dy: number): void;
  draw( g: CanvasRenderingContext2D, x: number, y: number ): void;
  update( x: number, y: number )
}

class FallStrategy {
  constructor(private falling: boolean) { };
  isFalling() { return this.falling; }
  drop(tile: Tile, x: number, y: number): void {
    if ( this.falling ) {
      map[y + 1][x] = tile;   // "fall"
      map[y][x] = new Air();  // replace previous spot with air
    }
  }
  update( tile: Tile, x: number, y: number ): void {
    this.falling = map[y + 1][x].isAir(); // set the falling flag
    this.drop(tile, x, y);
  }
}

class Air implements Tile {
  isAir() { return true; }
  isLock1() { return false; }
  isLock2() { return false; }

  moveHorizontal(dx: number) {
    moveToTile(playerx + dx, playery);
  }
  moveVertical(dy: number) {
    moveToTile(playerx, playery + dy);
  }
  draw( g: CanvasRenderingContext2D, x: number, y: number ) { }
  update( x: number, y: number ) { }
}

class Flux implements Tile {
  isAir() { return false; }
  isLock1() { return false; }
  isLock2() { return false; }

  moveHorizontal(dx: number) {
    moveToTile(playerx + dx, playery);
  }
  moveVertical(dy: number) {
    moveToTile(playerx, playery + dy);
  }
  draw( g: CanvasRenderingContext2D, x: number, y: number ) {
    g.fillStyle = "#ccffcc";
    g.fillRect(x * TILE_SIZE, y * TILE_SIZE, TILE_SIZE, TILE_SIZE);
  }
  update( x: number, y: number ) { }
}

class Unbreakable implements Tile {
  isAir() { return false; }
  isLock1() { return false; }
  isLock2() { return false; }

  moveHorizontal(dx: number) { }
  moveVertical(dy: number) { }
  draw( g: CanvasRenderingContext2D, x: number, y: number ) {
    g.fillStyle = "#999999";
    g.fillRect(x * TILE_SIZE, y * TILE_SIZE, TILE_SIZE, TILE_SIZE);
  }
  update( x: number, y: number ) { }
}

class Player implements Tile {
  isAir() { return false; }
  isLock1() { return false; }
  isLock2() { return false; }

  moveHorizontal(dx: number) { }
  moveVertical(dy: number) { }
  draw( g: CanvasRenderingContext2D, x: number, y: number ) {
    g.fillStyle = "#ff0000";
    g.fillRect(x * TILE_SIZE, y * TILE_SIZE, TILE_SIZE, TILE_SIZE);
  }
  update( x: number, y: number ) { }
}

class Stone implements Tile {
  private fallStrategy: FallStrategy;
  constructor(falling: boolean) {
    this.fallStrategy = new FallStrategy( falling );
  }
  isAir() { return false; }
  isLock1() { return false; }
  isLock2() { return false; }

  moveHorizontal(dx: number) {
    if ( ! this.fallStrategy.isFalling()
        && map[playery][playerx + dx + dx].isAir()
        && ! map[playery + 1][playerx + dx].isAir()) {
      map[playery][playerx + dx + dx] = map[playery][playerx + dx];
      moveToTile(playerx + dx, playery);
    }
  }
  moveVertical(dy: number) { }
  draw( g: CanvasRenderingContext2D, x: number, y: number ) {
    g.fillStyle = "#0000cc";
    g.fillRect(x * TILE_SIZE, y * TILE_SIZE, TILE_SIZE, TILE_SIZE);
  }
  update( x: number, y: number ) {
    this.fallStrategy.update( this, x, y );
  }
}

class Box implements Tile {
  private fallStrategy: FallStrategy;
  constructor(falling: boolean) {
    this.fallStrategy = new FallStrategy( falling );
  }
  isAir() { return false; }
  isLock1() { return false; }
  isLock2() { return false; }

  moveHorizontal(dx: number) {
    if ( ! this.fallStrategy.isFalling() && ! map[playery + 1][playerx + dx].isAir() ) {
      map[playery][playerx + dx + dx] = map[playery][playerx + dx];
      moveToTile(playerx + dx, playery);
    }
  }
  moveVertical(dy: number) { }
  draw( g: CanvasRenderingContext2D, x: number, y: number ) {
    g.fillStyle = "#8b4513";
    g.fillRect(x * TILE_SIZE, y * TILE_SIZE, TILE_SIZE, TILE_SIZE);
  }
  update( x: number, y: number ) {
    this.fallStrategy.update( this, x, y );
  }
}

class Key1 implements Tile {
  isAir() { return false; }
  isLock1() { return false; }
  isLock2() { return false; }

  moveHorizontal(dx: number) {
    removeLock1();
    moveToTile(playerx + dx, playery);
  }
  moveVertical(dy: number) {
    removeLock1();
    moveToTile(playerx, playery + dy);
  }
  draw( g: CanvasRenderingContext2D, x: number, y: number ) {
    g.fillStyle = "#ffcc00";
    g.fillRect(x * TILE_SIZE, y * TILE_SIZE, TILE_SIZE, TILE_SIZE);
  }
  update( x: number, y: number ) { }
}

class Lock1 implements Tile {
  isAir() { return false; }
  isLock1() { return true; }
  isLock2() { return false; }

  moveHorizontal(dx: number) { }
  moveVertical(dy: number) { }
  draw( g: CanvasRenderingContext2D, x: number, y: number ) {
    g.fillStyle = "#ffcc00";
    g.fillRect(x * TILE_SIZE, y * TILE_SIZE, TILE_SIZE, TILE_SIZE);
  }
  update( x: number, y: number ) { }
}

class Key2 implements Tile {
  isAir() { return false; }
  isLock1() { return false; }
  isLock2() { return false; }

  moveHorizontal(dx: number) {
    removeLock2();
    moveToTile(playerx + dx, playery);
  }
  moveVertical(dy: number) {
    removeLock2();
    moveToTile(playerx, playery + dy);
  }
  draw( g: CanvasRenderingContext2D, x: number, y: number ) {
    g.fillStyle = "#00ccff";
    g.fillRect(x * TILE_SIZE, y * TILE_SIZE, TILE_SIZE, TILE_SIZE);
  }
  update( x: number, y: number ) { }
}

class Lock2 implements Tile {
  isAir() { return false; }
  isLock1() { return false; }
  isLock2() { return true; }

  moveHorizontal(dx: number) { }
  moveVertical(dy: number) { }
  draw( g: CanvasRenderingContext2D, x: number, y: number ) {
    g.fillStyle = "#00ccff";
    g.fillRect(x * TILE_SIZE, y * TILE_SIZE, TILE_SIZE, TILE_SIZE);
  }
  update( x: number, y: number ) { }
}

function drawPlayer( g: CanvasRenderingContext2D ) {
  g.fillStyle = "#ff0000";
  g.fillRect(playerx * TILE_SIZE, playery * TILE_SIZE, TILE_SIZE, TILE_SIZE);
}

function gameLoop() {
  let before = Date.now();
  update();
  draw();
  let after = Date.now();
  let frameTime = after - before;
  let sleep = SLEEP - frameTime;
  setTimeout(gameLoop, sleep);
}

window.onload = () => {
  transformMap();
  gameLoop();
}

const LEFT_KEY = 37;
const UP_KEY = 38;
const RIGHT_KEY = 39;
const DOWN_KEY = 40;
window.addEventListener("keydown", e => {
  if (e.keyCode === LEFT_KEY || e.key === "a") inputs.push( new Left() );
  else if (e.keyCode === UP_KEY || e.key === "w") inputs.push( new Up() );
  else if (e.keyCode === RIGHT_KEY || e.key === "d") inputs.push( new Right() );
  else if (e.keyCode === DOWN_KEY || e.key === "s") inputs.push( new Down() );
});

