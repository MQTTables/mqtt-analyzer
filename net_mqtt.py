import json
import scapy.contrib.mqtt as mqtt
from scapy.all import *
from net_json import to_json

from pprint import pprint

packets = rdpcap('mqtt_packets_tcpdump.pcap')

for i, packet in enumerate(packets):
    if packet.haslayer(mqtt.MQTT):
        jrepr = to_json(packet)
        pprint(jrepr)
        with open(f'json/p_{str(i)}.json', 'w') as write_file:
            json.dump(jrepr, write_file, indent=4, separators=(',', ': '))