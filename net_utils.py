def mqtt_type(keys):
    if 'mqtt_connect' in keys:
        return 'connect'
    elif 'mqtt_connack' in keys:
        return 'connack'
    elif 'mqtt_publish' in keys:
        return 'publish'
    elif 'mqtt_puback' in keys:
        return 'puback'
    elif 'mqtt_pubrec' in keys:
        return 'pubrec'
    elif 'mqtt_pubrel' in keys:
        return 'pubrel'
    elif 'mqtt_pubcomp' in keys:
        return 'pubcomp'
    elif 'mqtt_subscribe' in keys:
        return 'subscribe'