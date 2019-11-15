'''
Project: MQTT Traffic Analyzer
Module: Packeda
- Packet grabber; calls database operator
'''

from scapy.all import sniff

import p_database


class Grabber:
    def __init__(self, db_name, t_name):
        self.db_name = db_name
        self.t_name = t_name
        self.db = p_database.Database(db_name, t_name)

    def start(self):
        while True:
            packet = sniff(count=1)[0]
            self.db.add_packets([packet])