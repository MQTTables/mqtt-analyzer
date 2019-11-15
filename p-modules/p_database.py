'''
Project: MQTT Traffic Analyzer
Module: Packeda
- Direct database operator
'''

import sys
import json
import os.path
import sqlite3
import scapy.contrib.mqtt as mqtt
from scapy.all import *
from pprint import pprint

from p_json import to_json, mqtt_type

class Database:
    def __init__(self, db_name, t_name):
        self.db_name = db_name
        self.t_name = t_name

        self.conn = sqlite3.connect(db_name + '.db')
        self.c = self.conn.cursor()

        self.c.execute(f'''CREATE TABLE IF NOT EXISTS "{t_name}" (
                    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
                    time_rel REAL NOT NULL,
                    ip_src TEXT NOT NULL,
                    ip_dst TEXT NOT NULL,
                    port_src INTEGER NOT NULL,
                    port_dst INTEGER NOT NULL,
                    mqtt_type TEXT);''')

        self.c.execute(f'''CREATE TABLE IF NOT EXISTS "{t_name}_data" (
                    id INTEGER PRIMARY KEY NOT NULL,
                    json TEXT);''')

    def add_packets(self, packets):
        time_ref = float(packets[0].time)
        i = 0

        for packet in packets:
            if packet.haslayer(mqtt.MQTT):
                jrepr = to_json(packet)
                jrepr['time']['time_abs'] = float(packet.time)
                jrepr['time']['time_rel'] = float(packet.time) - time_ref
                # pprint(jrepr)
                # with open(f'json/p_{str(i)}.json', 'w') as write_file:
                #    json.dump(jrepr, write_file, indent=4, separators=(',', ': '))
                
                entry = [
                    jrepr['time']['time_rel'],
                    jrepr['ip']['src'],
                    jrepr['ip']['dst'],
                    jrepr['tcp']['sport'],
                    jrepr['tcp']['dport'],
                    mqtt_type(jrepr['mqtt_fixed_header']['type'])
                ]
                self.c.execute(f'''INSERT INTO "{self.t_name}"" (time_rel, ip_src, ip_dst, port_src, port_dst, mqtt_type)
                            VALUES(?, ?, ?, ?, ?, ?)''', entry)
                self.c.execute(f'''INSERT INTO "{self.t_name}_data" (id, json)
                    VALUES(?, ?)''', [i, json.dumps(jrepr)])

                self.conn.commit()
                i += 1
        return