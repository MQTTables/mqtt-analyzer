import sys
import scapy.contrib.mqtt as mqtt
from scapy.all import *

packets = rdpcap(input())

for packet in packets:
    print(packet.summary())
    print(packet.show())
    '''if packet.haslayer(mqtt.MQTT):
        print('MQTT')'''