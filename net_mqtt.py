import sys
import json
import os.path
import sqlite3
import scapy.contrib.mqtt as mqtt
from scapy.all import *
from pprint import pprint

from net_json import to_json, mqtt_type

class Database:
    def __init__(self, db_name, t_name):
        conn = sqlite3.connect(db_name + '.db')
        c = conn.cursor()

        c.execute(f'''CREATE TABLE IF NOT EXISTS {t_name}(
                    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
                    time_rel REAL NOT NULL,
                    ip_src TEXT NOT NULL,
                    ip_dst TEXT NOT NULL,
                    port_src INTEGER NOT NULL,
                    port_dst INTEGER NOT NULL,
                    mqtt_type TEXT);''')

        c.execute(f'''CREATE TABLE IF NOT EXISTS {t_name}_data(
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
                c.execute(f'''INSERT INTO {sys.argv[3]} (time_rel, ip_src, ip_dst, port_src, port_dst, mqtt_type)
                            VALUES(?, ?, ?, ?, ?, ?)''', entry)
                c.execute(f'''INSERT INTO {sys.argv[3]}_data (id, json)
                    VALUES(?, ?)''', [i, json.dumps(jrepr)])

                conn.commit()
                i += 1
        return