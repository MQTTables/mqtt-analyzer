import struct
import socket


def eth_addr(raw):
    toi = lambda a: '{:02x}'.format(raw[a])
    return f'{toi(0)}:{toi(1)}:{toi(2)}:{toi(3)}:{toi(4)}:{toi(5)}'

def net_decode(packet):
    eth_length = 14

    eth_header = packet[0:eth_length]
    eth = struct.unpack('!6s6sH', eth_header)
    eth_protocol = socket.ntohs(eth[2]) # [- Protocol ()

    mac_src = eth_addr(packet[6:12])
    mac_dst = eth_addr(packet[0:6])

    return mac_src, mac_dst