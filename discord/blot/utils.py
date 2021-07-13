import requests

def ensureGameExists(guildId):
    r = requests.get("http://localhost:8080/games")
    if not(guildId in r.json()):
        r = requests.put("http://localhost:8080/games/" + guildId)

def getBattleStateMessage():
    return "Okay... let's see how the war is waging"
