import sys
import json
import os.path
import sqlite3
import scapy.contrib.mqtt as mqtt
from scapy.all import *
from pprint import pprint

from net_json import to_json, mqtt_type

print(sys.argv)

conn = sqlite3.connect(sys.argv[2] + '.db')
c = conn.cursor()

c.execute(f'''CREATE TABLE IF NOT EXISTS {sys.argv[3]}(
              id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
              time_rel REAL NOT NULL,
              ip_src TEXT NOT NULL,
              ip_dst TEXT NOT NULL,
              port_src INTEGER NOT NULL,
              port_dst INTEGER NOT NULL,
              mqtt_type TEXT);''')

c.execute(f'''CREATE TABLE IF NOT EXISTS {sys.argv[3]}_data(
              id INTEGER PRIMARY KEY NOT NULL,
              json TEXT);''')

packets = rdpcap(sys.argv[1])


time_ref = float(packets[0].time)
i = 0

for packet in packets:
    if packet.haslayer(mqtt.MQTT):
        jrepr = to_json(packet)
        jrepr['time']['time_abs'] = float(packet.time)
        jrepr['time']['time_rel'] = float(packet.time) - time_ref
        # pprint(jrepr)
        with open(f'json/p_{str(i)}.json', 'w') as write_file:
            json.dump(jrepr, write_file, indent=4, separators=(',', ': '))
        
        mtype = mqtt_type(jrepr.keys())
        
        entry = [
            jrepr['time']['time_rel'],
            jrepr['ip']['src'],
            jrepr['ip']['dst'],
            jrepr['tcp']['sport'],
            jrepr['tcp']['dport'],
            mtype
        ]
        c.execute(f'''INSERT INTO {sys.argv[3]} (time_rel, ip_src, ip_dst, port_src, port_dst, mqtt_type)
                      VALUES(?, ?, ?, ?, ?, ?)''', entry)
        if mtype:
            c.execute(f'''INSERT INTO {sys.argv[3]}_data (id, json)
                        VALUES(?, ?)''', [i, json.dumps(jrepr)])

        conn.commit()
        i += 1