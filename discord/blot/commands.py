from flask import jsonify

import requests
import webcolors

## Commands
def bloop(data):
    x = 0
    y = 0
    for opt in data['data']['options'][0]['options']:
        if opt['name'] == 'x':
            x = opt['value']
        elif opt['name'] == 'y':
            y = opt['value']
        
    guildId = data['guild_id']
    user = data['member']['user']['id']
    playerConfig = {
        "start": {
            "x": x,
            "y": y,
        }
    }
    r= requests.post("http://localhost:8080/playerConfigs/" + guildId + "/" + user, json=playerConfig)
    print(r.text)
    return jsonify({
        "type": 4,
        "data": {
            "tts": False,
            "content": "Set location for user " + data['member']['user']['username'],
            "embeds": [],
            "allowed_mentions": { "parse": [] }
        }
    })

def ploint(data):
    guildId = data['guild_id']
    user = data['member']['user']['id']
    playerConfig = {
        "priority": {
            "v": data['data']['options'][0]['options'][0]['value']
        }
    }
    r= requests.post("http://localhost:8080/playerConfigs/" + guildId + "/" + user, json=playerConfig)
    print(r.text)
    return jsonify({
        "type": 4,
        "data": {
            "tts": False,
            "content": "Set clolour for user " + data['member']['user']['username'],
            "embeds": [],
            "allowed_mentions": { "parse": [] }
        }
    })

def colorToRGB(name):
    (r,g,b) = webcolors.name_to_rgb(name, spec='css3')
    return {
        "R":r,
        "G":g,
        "B":b,
    }


def clolour(data):
    guildId = data['guild_id']
    user = data['member']['user']['id']
    color = data['data']['options'][0]['options'][0]['value']
    playerConfig = {
        "color": colorToRGB(data['data']['options'][0]['options'][0]['value'])
    }
    print(playerConfig)
    r= requests.post("http://localhost:8080/playerConfigs/" + guildId + "/" + user, json=playerConfig)
    print(r.text)
    return jsonify({
        "type": 4,
        "data": {
            "tts": False,
            "content": "Set clolour for user " + data['member']['user']['username'],
            "embeds": [],
            "allowed_mentions": { "parse": [] }
        }
    })


