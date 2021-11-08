int numberOfNumbers = 200;
int [] values = new int[numberOfNumbers];
int deltaX = 1;

void setup() {
  size(1000, 500);
  deltaX = width/numberOfNumbers;

  for (int i = 0; i < values.length; i++) {
    int randomValue = int(random(0, height));
    values[i] = randomValue;
  }
}

void draw() {
  background(255);
  stroke(180);
  for (int i = 0; i < Math.min(frameCount, values.length); i++) {
    fill(values[i], 255, numberOfNumbers);
    rect(i*deltaX, height-values[i], deltaX, values[i]);
  }
}
