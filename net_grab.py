from scapy.all import sniff

import net_mqtt


class Grabber:
    def __init__(self, db_name, t_name):
        self.db_name = db_name
        self.t_name = t_name
        self.db = net_mqtt.Database(db_name, t_name)

    def start(self):
        while True:
            packet = sniff(count=1)[0]
            self.db.add_packets([packet])