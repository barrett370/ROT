int ledPin = LED_BUILTIN;                // choose the pin for the LED
int inputPin = 2;               // choose the input pin (for PIR sensor)
int pirState = LOW;             // we start, assuming no motion detected
int val = 0;                    // variable for reading the pin status
 
void setup() {
  pinMode(ledPin, OUTPUT);      // declare LED as output
  pinMode(inputPin, INPUT);     // declare sensor as input
  digitalWrite(ledPin, LOW);
  Serial.begin(9600);
}
 
void loop(){
//  Serial.println("Loop");
  val = digitalRead(inputPin);  // read input value
//  Serial.println(val);
  if (val == HIGH)  // check if the input is HIGH
  {            
    digitalWrite(ledPin, HIGH);  // turn LED ON
  
    if (pirState == LOW) 
  {
      Serial.println("1"); // print on output change
      pirState = HIGH;
    }
  } 
  else 
  {
    digitalWrite(ledPin, LOW); // turn LED OFF
  
    if (pirState == HIGH)
  {
      Serial.println("0"); // print on output change
      pirState = LOW;
    }
  }
}

