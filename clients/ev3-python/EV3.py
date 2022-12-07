import time
import ev3api;
from enum import Enum;

from pprint import pprint

from ev3api.apis.tags import motor_api;
from ev3api.apis.tags import power_api;
from ev3api.apis.tags import sensor_api;
from ev3api.apis.tags import sound_api;

from ev3api import api_client;
from ev3api import exceptions;

from ev3api.models import Text;
from ev3api.models import Tone;

class EV3:
    
    class Motors(Enum):
        A = "A"
        B = "B"
        C = "C"
        D = "D"
        
        def __init__(self, name):
            self.name = name;
        
        def __str__(self):
            return self.name;
    
    
    
    def __init__(self, hostAdress):
        self.hostAdress = hostAdress;
        self.apiclient = api_client;
        # TODO: api_client.setBasePath()
        self.motorApi = motor_api(_apiclient);
        self.powerApi = power_api(_apiclient);
        self.sensorApi = sensor_api(_apiclient);
        self.soundApi = sound_api(_apiclient);
        
    
    """
    This method returns the ip-adress from the EV3
    @return the ip-adress   
    """
    def getHost_adress(self):
        return self.hostAdress;
    
    
    """
    The EV3 will do a beep sound.
    """
    def beep(self):
        try:
            self.soundApi.SoundBeepPost();
        except exceptions:
            print("TODO");
            
    """
    * The EV3 will play a tone.
    * @param frequency the specific frequenz for the tone
    * @param lengthMs the specific duration of the tone
    """
    def play_tone(self, frequenz, lengthMs):
        try:
            self.soundapi.SoundTonePost(Tone.frequency(frequenz).lengthMs(lengthMs)); # TODO
        except exceptions:
            print("TODO");
            
            
    """
    The EV3 will speak a specific text.
    @param text the spoken text for the EV3
    """
    def speak(self, text):
        try:
           self.soundapi.soundSpeakPost(Text.text(text)) #TODO
        except exceptions:
            print("TODO");
            
    """
    This method always returns immediately, whether or not the battery voltage level exists.
    @return the battery voltage level.
    """
    def voltage(self):
        try:
            return self.powerApi.PowerGet(); #TODO
        except exceptions:
             print("TODO");
        return -1;
    
    """
    This method always returns immediately, whether or not the battery current level exists.
    @return the battery current level.
    """     
    def current():
        #TODO
        return -1;
        
        
    """
    This method always returns immediately, whether or not the maximal battery voltage exists.
    @return the maximal battery voltage
    """
    def max_voltage():
        #TODO
        return -1         
    
    
    """
    This method always returns immediately, whether or not the minimal battery voltage exists.
    @return the minimal battery voltage.
    """
    def min_voltage():
        #TODO
        return -1         
    
    
    """
    This method always returns immediately, whether or not the battery technology description exists.
    @return the battery technology description
    """
    def technology(self):
        try:
            return self.apiclient.PowerGet();#TODO
        except exceptions:
            print("TODO");
        
        return None;
    
    """
    This method returns an flag, if the button is pressed or not
    @return the boolean if pressed or not
    """ 
    def button():
        #TODO
        return False;
    
    """
    The EV3 will flash the LEDs immediately.
    """
    def flash():
        #TODO
        return None;    
    
    
    """
    This method will set the LEDs of from the EV3
    """
    def led():
        #TODO
        return None;
    
    """
    This method will switch off the LEDs of from the EV3.
    """
    def led_off():
        #TODO
        return None;
    
    """
    This method will turn on the monitor from the EV3
    """
    def monitor_on():
        #TODO
        return None;
    
    """
    This method will turn off the monitor from the EV3
    """
    def monitor_off():
        #TODO
        return None;