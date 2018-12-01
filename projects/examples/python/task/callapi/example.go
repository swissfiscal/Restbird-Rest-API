# encoding: utf-8
import requests
from restapi import RestApi
from restBirdCtx import RestBirdCtx

"""
write any your Python scripts here
"""

#Call Demo_Rest/Rest_Python api0 with env Demo_Rest_Env
sess = requests.Session()
ctx  = RestBirdCtx(setting.local, "Demo_Rest_Env", sess, "")

api = RestApi("Demo_Rest/Rest_Python", "api0", ctx)
api.runApi()

#Get/Set global env
ctx.setGlobal("hello", "world")
v = ctx.getGlobal("hello")
print (v)

#Get Env
print(ctx.loadEnvVaribles("Demo_Rest_Env"))
print(ctx.vars["counter"])
