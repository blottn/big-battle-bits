from flask import jsonify

import json
import random
import requests
import webcolors

from utils import *


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
            "content": "Set location for user " + data['member']['user']['username'] + " to (" + str(x) + ", " + str(y) + ")",
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
            "content": "Set direction for user " + data['member']['user']['username'] + " to " + data['data']['options'][0]['options'][0]['value'] + " degrees",
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
        "A":255,
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
            "content": "Set clolour for user " + data['member']['user']['username'] + " to: " + playerConfig["color"],
            "embeds": [],
            "allowed_mentions": { "parse": [] }
        }
    })

def getPlayerConfig(data):
    guildId = data['guild_id']
    user = data['member']['user']['id']

    r = requests.get("http://localhost:8080/playerConfigs/" + guildId)
    return jsonify({
        "type": 4,
        "data": {
                "tts": False,
                "content": data['member']['user']['username'] + " config: ```" + json.dumps(r.json().get(user, {})) + "```",
                "embeds": [],
                "allowed_mentions": {"parse": []}
            }
        })
    userConfig = r.json()[user]

def forceStep(data):
    guildId = data['guild_id']
    iters = 5
    cmd = data['data']['options'][0]
    if 'options' in cmd:
        iters = cmd['options'][0]['value']
    
    print("Stepping with " + str(iters) + " iterations")

    for i in range(0,iters):
        r = requests.get("http://localhost:8080/games/" + guildId + "/step")
        print(r.text)
    return jsonify({
        "type": 4,
        "data": {
            "tts": False,
            "content": getBattleStateMessage(),
            "embeds":[
                {
                    "type": "image",
                    "image": {
                        "url": "https://blot.blottn.ie/battlegrounds/" + str(guildId) + "/" + str(random.randint(0,100000)),
                    }
                }
            ],
            "allowed_mentions": {"parse": []}
        }
    })

def getState(data):
    guildId = data['guild_id']

    return jsonify({
        "type": 4,
        "data": {
            "tts": False,
            "content": getBattleStateMessage(),
            "embeds":[
                {
                    "type": "image",
                    "image": {
                        "url": "https://blot.blottn.ie/battlegrounds/" + str(guildId) + "/" + str(random.randint(0,100000)),
                    }
                }
            ],
            "allowed_mentions": {"parse": []}
            }
        })

def start(data):
    guildId = data['guild_id']
    r = requests.get("http://localhost:8080/start/" + guildId)
    if r.status_code >= 400:
        return jsonify({"type": 4, "data":{"content":"Error starting game: " + r.text}})

    return jsonify({
        "type": 4,
        "data": {
            "content": "Started game for " + guildId + ", updates to location, color. But you are still free to change priority direction."
        }
    })

def reset(data):
    guildId = data['guild_id']
    r = requests.get("http://localhost:8080/reset/" + guildId)
    print(r.text)
    return jsonify({
        "type": 4,
        "data": {
                "content": "Reset game for " + guildId,
            }
        })
