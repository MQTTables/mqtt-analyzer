import argparse
from scapy.all import rdpcap

import net_mqtt
import net_grab


parser = argparse.ArgumentParser(description='Process MQTT packets')
parser.add_argument('db_name', metavar='D', type=str, help='database name to operate')
parser.add_argument('t_name', metavar='T', type=str, help='table name to operate')
parser.add_argument('mode', metavar='M', type=str, help='operation mode (pcap / grab)')
parser.add_argument('f_name', metavar='F', type=str, help='.pcap file path (if any)')

args = parser.parse_args()
if args.mode == 'pcap':
    db = net_mqtt.Database(args.db_name, args.t_name)
    db.add_packets(rdpcap(args.f_name))
elif args.mode == 'grab':
    gr = net_grab.Grabber(args.db_name, args.t_name)