// Based off of NOC_I_5_AcceptReject.
int[] randomValues;
void setup() {
  size(800, 200);
  randomValues = new int[20];
  for(int i = 0; i < 20; i++) {
    randomValues[i] = (int)random(0,height);
  }
}

void draw() {
  stroke(0);
  strokeWeight(2);
  fill(127);
  int w = width/randomValues.length;
  for(int x = 0; x < randomValues.length; x++) {
    rect(
      x*w,                    // x
      height-randomValues[x], // y
      w-1,                    // width
     randomValues[x]);        // height
  }
}
