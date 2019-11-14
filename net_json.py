import json
from scapy.base_classes import BasePacket, Gen, SetGen, Packet_metaclass, \
    _CanvasDumpExtended
from scapy.fields import StrField, ConditionalField, Emph, PacketListField, \
    BitField, MultiEnumField, EnumField, FlagsField, MultipleTypeField, FlagValue

def to_json(packet):
    d = {}
    d['time'] = {}
    for i in range(100):
        try:
            layer = packet[i]
        except IndexError:
            break
        dl = {}
        for f in layer.fields_desc:
            if isinstance(f, ConditionalField) and not f._evalcond(layer):
                continue

            fvalue = layer.getfieldval(f.name)
            if type(fvalue) == FlagValue:
                fvalue = str(fvalue)

            if type(fvalue) == bytes:
                fvalue = fvalue.decode()

            dl[f.name] = fvalue
            # print(f'{f.name}: {fvalue}')
            '''if isinstance(fvalue, Packet) or (f.islist and f.holds_packets and isinstance(fvalue, list)):
                fvalue_gen = SetGen(fvalue, iterpacket=0)
                for fvalue in fvalue_gen:
                    dl[]'''
        d[packet[i].name.lower().replace(' ', '_')] = dl
    d['tcp']['options'] = []
    return d