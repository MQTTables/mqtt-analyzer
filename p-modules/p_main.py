'''
Project: MQTT Traffic Analyzer
Module: Packeda
- Main script; entry point for external calls
'''

import argparse
from scapy.all import rdpcap

import p_database
import p_grabber


parser = argparse.ArgumentParser(description='Process MQTT packets')
parser.add_argument('db_name', metavar='D', type=str, help='database name to operate')
parser.add_argument('t_name', metavar='T', type=str, help='table name to operate')
parser.add_argument('mode', metavar='M', type=str, help='operation mode (pcap / grab)')

args = parser.parse_args()
print(args)
if args.mode == 'pcap':
    db = p_database.Database(args.db_name, args.t_name)
    db.add_packets(rdpcap('.cache/' + args.t_name + '.pcap'))
elif args.mode == 'grab':
    gr = p_grabber.Grabber(args.db_name, args.t_name)
    gr.start()