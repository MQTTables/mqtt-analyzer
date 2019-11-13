import socket
import struct
import time

import net_packet

s = socket.socket(socket.AF_INET, socket.SOCK_RAW, socket.IPPROTO_TCP)
c = 0
time_s = time.time()

while True:
    packet = s.recvfrom(65565)[0]
    time_c = time.time()

    mac_src, mac_dst = net_packet.net_decode(packet)
    
    print(f'#{c} - {time_c}\nSource: {mac_src}\nDestination: {mac_dst}\n')
    c += 1