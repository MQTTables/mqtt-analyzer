import socket
import struct
import time

s = socket.socket(socket.AF_INET, socket.SOCK_RAW, socket.IPPROTO_TCP)
c = 0
time_s = time.time()

while True:
    packet = s.recvfrom(65565)
    time_c = time.time()

    packet = packet[0]

    ip_header = struct.unpack('!BBHHHBBH4s4s', packet[0:20])

    magic_verihl = ip_header[0]
    version = magic_verihl >> 4
    ihl = magic_verihl & 0xF

    iph_length = ihl * 4

    ttl = ip_header[5]
    protocol = ip_header[6]
    addr_s = socket.inet_ntoa(ip_header[8])
    addr_d = socket.inet_ntoa(ip_header[9])

    print(f'''[#{c}, {time_c - time_s}]
    Version: {version}
    IHL: {ihl}
    TTL: {ttl}
    Protocol: {protocol}
    Source Address: {addr_s}
    Destination Address: {addr_d}''')

    c += 1