//  ResourceData.go  -  Generated resource file.
//  Resources - E.B.Smith  -  Sat Mar 19 10:47:05 PDT 2016


package ApplePushService


import (
    "bytes"
    "strings"
    "reflect"
    "io/ioutil"
    "compress/gzip"
    "encoding/base64"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
)


type ResourceData struct {
    name    string
    data    string
    header  gzip.Header
}


var Resource ResourceData = ResourceData{}


func (resource *ResourceData) Name() string {
    return resource.name
}


func (resource *ResourceData) Header() gzip.Header {
    return resource.header
}


func (resource *ResourceData) Bytes() *[]byte {
    compressedBytes, error := base64.StdEncoding.DecodeString(resource.data)
    if error != nil { return nil; }

    buffer := bytes.NewBuffer(compressedBytes)
    gz, error := gzip.NewReader(buffer)
    if error != nil { Log.LogError(error); return nil; }
    defer gz.Close()

    rawbytes, error := ioutil.ReadAll(gz)
    if error != nil { Log.LogError(error); return nil; }
    resource.header = gz.Header

    return &rawbytes;
}


func (resource *ResourceData) ResourceBytesNamed(name string) *[]byte {
    name = Util.ReplaceCharactersNotInSetWithRune(name, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_", '_')
    name = strings.Title(name)
    Log.Debugf("Loading resource '%s'.", name)
    methodV := reflect.ValueOf(resource).MethodByName(name)
    if methodV.IsValid() {
        values := methodV.Call([]reflect.Value{ })
        //Log.Debugf("Values: %+v.", values)
        if len(values) == 1 && values[0].CanInterface() {
            data, found := values[0].Interface().(*ResourceData)
            if (found) { return data.Bytes() }
        }
    }
    return nil
}


func (resource *ResourceData) Com_blitzhere_blitzhere_labs_development_key() *ResourceData {
    return &ResourceData{
    name: "Com_blitzhere_blitzhere_labs_development_key",
    data:
`
H4sIAJmQ7VYCA21Vxw6rVhDd8xXsUWKqgackEp1LMwZM29HBYDo28PVx3lski8xi
NEejGR0dTeGTCubWdW7SbS0WCP5aOTdFn3eHlbyKH3A2vH5Pu2Y962Iu/o1+65J0
+T0v3kU3jK+iX3+WdkOWdHpxAPEHTBIwysASDkskLFxhUoBRESZYmKD/gSgHkxTM
fL0MyzyMXWFJgHEBlgmYkWHo2+M/tH7Af1jDf/Bf0G//GC8pwIIdl4NtB/icJ8G6
FP3MQCYA0sgBnuN0gbtLHObl4oSq+c1MVMv12Disq5sSYGIn2U8quqBoaOWTRbia
cTkte4HsMwzYqc4UGuESc0WXwEz4y8UP87ShLZEPRD+rhM/gJqojIzGLitaivB/v
63v3VMtCoV3jT5OhEA3Xh6x0GLC/maguT0FIE6yOTNNXgLSP15y/PCWcZO/kcgFX
7H0NJXx3agQarNIN9ImKp9dTKK11HK7P601UWL8iRc8Hlvd8MR2Sit7lMamNsh61
QHcpIpjTCEqZgh5ddh3vrkuE7V75TtLqnsWCgVJAf9G2zn41ITNrB5VE3SHrIDRu
hdPESbt88haZ0hekoQKW0Zhlsid5kRJ+a3jwOfHbnQwa0dQYXrWr6g5E7s7x3PAV
W+3n8rOGb3TZyPmGQUz9ng6gh155qXa5ehIYZmk5TdpSDAbVzHrK1tpJuu6vZm5i
umHkOtlzMTUWO5auGwOxKMjcR+EXapW4c3sgBr9oO6lqPNVPmrHoBAgYou0e13LG
0F23OMIrEMaeTIKq8oWAqtbyvQXFZpbobqUwZ0UqDXVW94utC+q9ajoC6VObdpVr
t78cnn5sw2KchqfUkf50aeg4ahVH1iU2pNhQJ55CbPflX/yckVoF09sL4VFFMt0d
1Auc+isKRUTuKda1Ri2sMy4Q0+l5yZ8hVwnHca6WJGE8rb3pm/doeMwIX9jDugPk
GTaYMBRvkszr01/ZE5GYmV5XChJAv9EAFapI4pC4Hx2As6h0CQJxRh/utnbWYXxu
482lbRqPhI8LsNI4bcwwnrcPKAgIy6QgPd5cPe67MjnmGLo2BwpxuHWBGFqRWriF
hrK8+5HPDHOGmssEZi7TpOrIb50AZYtnAIqxJkILhNaPtg0XZq3s43yebiAjvBvS
+Kr2zM2JSYsYaa1CognmJ2N8ut4hju0vZHA1aFMc/KxWjMgSpiUd7zn/7Cq26Zv9
rMSQeZM2rXuRX5IJSGlLMnHWcvbugLZq7QSQnjFqn3LoSITBxfJo7HRXtv4joDqi
zzA2M16Zt7nF2tORWVUNXiQ6qFdm96DN4P3TCIcNK6u622WjU611SvOEOZk4Ot0N
797G5RdlwiHw+0jjxxzZ1Tvq7bbLoeTjIlGefPRix9M2NSm5HW8UUaOfa+721IfE
AINRG0d67CFgQL4An1jyOHVJZKq/awAtlUByS0rFi3n38pbgMzZ6oApfe4fwnHqj
fnviZD+z16MoKD9e13HF9epeAUXhgZdGkGeE2DwR7ygJuelyUyfR66zg11xgJ+Id
ZYJUxwxw3o9fRHSUERCkO+NktT4sJBdDfnU5GPV+VjK7nF5pKEEenOp7xhNvmFmW
m+qiM45NeXTGSXUk8Sq/x5CszccwkngRn9DCqM/tEB70OEnkEHyUHhPzjh1wB80b
tCFoJuXpoP9g61fWyl5GLRuG4JTpAsxh/pqhiPRJl2u/lHmKWa2z2ySeyxWi9j+9
WD5rTjRwtdysM0VOsW9HSq3S283zqAHp5nVLIb15FZ7cmK6Wm8GUPmyumsDiuTor
fT5e6znqvHPWqEnfScsvTvThdTqmtUHlj5dQexnUvNks0J7WmjM3ThZf3ileGy8+
3whq2pWhHPk6Yolqcr3eRYjMiHiiDNWff/56OJIl/v+7+Rs3f7PrRgcAAA==
`   }
}



func (resource *ResourceData) Com_blitzhere_blitzhere_labs_development_pem() *ResourceData {
    return &ResourceData{
    name: "Com_blitzhere_blitzhere_labs_development_pem",
    data:
`
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
+f35lgzx3x9v+M9fsf8BU8Bb45oJAAA=
`   }
}



func (resource *ResourceData) Com_blitzhere_blitzhere_labs_production_key() *ResourceData {
    return &ResourceData{
    name: "Com_blitzhere_blitzhere_labs_production_key",
    data:
`
H4sIAJmQ7VYCA21Vx66sCA7d8xXsUTdFhqfRSORUUOS0I+cMVRRf3/d1L6YX45XD
sWUd6dhcWoPscWxtdh7lDoA/Vm1tORXD10zH8heYz+Of2dAed1Nu5f+8P4Y02/9c
trk486Odp787hzlPB738qsIvUJRAnAcRDCR4EKVAVAQ5EsQIkJZARgQR6bfzkxfp
3zAKBR/sb4xIgJQEAj8z/rXVL/A/5vyv+L/AH7+NE2XVBB2XBS1HDVhPBHUx/rsC
GKoqzh+VY1mdZ22R3XJYo6RtideMaHpJGA2Kdy2UNFZasx2exL9TqXPwcp3vppNy
CsBrOuy7hAky/8FTWf3kG9JU8L1cbLubL9p2sKd9BIjcUK2xDBY5D0RhZa5ANAqr
rCIgjZCctO5X03Ol5CnO1oibve62MiKllEWeghmIZBiRK2Is2bjIp+TiKT/ZobVk
tfuYwFDQCl9JqvSksMBvhs6czi8qFKTfnq7y2h91w7lPOCRDLH6nn30RVi4a6O9b
iKThMmLAqRp2RXmisW9GaGoibDk8S6M7mUWbeDOYjglU7gaKE4Z0LilPm5sqFIOH
3aPeBLnqAPwdNCwo4kBC2sn1XiNnh2NCz+Mw5Cwpe8wr2h8fVWBtlmPnH7KVcU4w
FE9UM37L1A5EhD0sSuXOlozTi6QGamkEq1kZl07oQiLIVn5B8gLLweGn9ahxGNxI
bquSQt+7nUYDn+3jeKuraRLuKwy03rC/auqwQBlhL5DCYEjqoz4m1ekpwbVzmZGB
FK9DxRQqMZ5BA+zFd/1a7084XVQivw11sPhuEr5UN09fjr13dn+4+0O0mKXj+WA6
e99qh/metvKF4YoMNF5YhKaNCSVvy7lRnQzuGdZVnyP9pDE8jZOJ1ZcTt8NrbDqk
togLovM9D7ue0eNRAz5WNQwN6rZ3qXXIOuCIM3rtaNiYiHnYKyyuvc0bg2FQX3mL
8nELL8H2fWklZrgq3A0Qv81rZkW+jkUWH4yDUWzyMDhkfVo+cmZ7dEVaGKkRGyZ2
ALfm6ULxtMbbO61g2/EnwNaSBM9Ee30UxFxrvtfq56yhh3C22dzU7LFlQtHHBPSs
cM/y0nbajamKHB6muqA9KIAysM28hLekIjBhqkLaQFhELx9bdT6aNn8fE+cJx3N6
ML2ETWS5myIqMcbfG1/adAMv8d4nv7sfVBxYiZ8TOZIyuEVZnItYPikX6aJ0fp/d
+CeK91I32r7Am9td9KoQD8sAfnSeya1JyMhVu+7ZlJSoFgH0nqu3g/t9PNd1o1Ii
2ia4B3/iqqX3SWjN6xVp6WvBdMCCYBjuIQSJXmlnHWSoTDasN7NXSg4VijZNvvDj
8w/J2Ilwc3af6xjS1vd+qkwaAY/jThBFMUtfw8w39PQj9UrD07V5DuYx4+7JSSSM
1YCyzi5cTzVJ7f1IjQISk/7ZNDOQEzyu2E/Ec9v6LN0634QsI6uwmsxAg5JbiHmP
ExH8Y99eXWaVzNtVPs/bj/pcVz9KQCg/2IN8zN6scdxmw1+vIVr9N8ss5VBO9NXU
nHtzcmpbfr+TM8pV7iWyFqmvTrH3PBBQm1T0z6Gyivyl8RoUYq3n8O6FUnGLJQ1U
yE55rbTXD6G2b5Lf5tnl/qgVZevuvUVAqeFm0Xj5h02lulVHmgpPToddm1CE4LU+
uO3nMj7x08a0pO2T/uMS20Or9C/0JIy0oQBs9BXW1rmaVTgzZU++1SG2ZDvau9bI
/ojB68Ri1Sf4nlR5cX+p7knNMxHjlt5na9kAhDH6N6FSp13Lvsdy8+Mtv45MKSFy
EeA3PLY+XfWnlo8GjpJ9KzumIogbfGVLmKN3BFRG/AkSJMuQABdNdHKvB76p3zDm
p7iLMm9YmwB+PPtFQ14Z2qsNQlz/PBvRFP7/q/kLAWfGd0EHAAA=
`   }
}



func (resource *ResourceData) Com_blitzhere_blitzhere_labs_production_pem() *ResourceData {
    return &ResourceData{
    name: "Com_blitzhere_blitzhere_labs_production_pem",
    data:
`
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
`   }
}



func (resource *ResourceData) Server_ca_cert_pem() *ResourceData {
    return &ResourceData{
    name: "Server_ca_cert_pem",
    data:
`
H4sIAJmQ7VYCA2VUubajOBDN+YrJffoYDNgQTCBAgNhs8dgzNoPZjBfWrx/cr5N+
o0Q6pbql0q1b9evXtgSoIOsfEdoOkpEIHPgx/iJMhKBeiSKY9QJMSAAFgmflku8k
YAlF/Sjrm8JPpACwKwNJFN4SdE0BKYByoViaBmH7LJl23pqJ7BAFeMYS2AmF5Qng
5Vp0RmcL2ySdvWSBRRot22QiL7mydTUrQJ5FoSeS1hrTDgxI0Vikak10YwpdmfvE
b8h0EV6xL9+MWz1rDqj+BHZUsaz0LzCfnZolEPSGTNWolMZD0nokgt/gyMezuYL5
G2Q6ev3Tz2rSzu6JqG2qMLD75MAUOPDIWOGXOLBZJBaLKWFGd0AuTyR7duBiVni2
VneyHDnebMvZATTx2yhBynS02EQKadp4gkUoeRjrcHrLf/1fkoEpMd8EzvjDEE0Y
7d+pGaSFXYr7/TpSmiE88EsqMkXY1UXaeq1xA/dEad5blivauCJCHw26AxtTNP9E
NlNdMXskwS3tuvhRpcL5BjeRBFdTggrxAYHZfPx0xAdvyZSmjX2r3PYxuQlCFtj3
hNb6TK0LXUKTJRX9R0gCqn6KBsoAnEWAOfC5Fwt9O0PwdCnjYevKa3hUCKSe3OIH
oZOjPU5v3QlP70KNm+iUMdgUVvwq75P1dnSeygc1Dq3obCg9xcH1flapgVZf+0YT
cE68QqWvggOzKm8D7KEoWcuzd0ENVJ5sdCUTxXa9gZEq6TzzUrqe6FMQt1/P0v5S
Gm8ICLMRxkt6O0bFWu2NA/OVyoebW0f7VLyPYVtFy17pTvOs+CJj5PVriYRD1A0u
8+AZqvUcInBWvzOMi44vrHG08etto7VwFy988q/WNqXhKxToILnyO1aUvaJ0QL8L
3qhXWoVZ3T2h3qMs7+7emHOgKt2bJ6SgFmNnBKwWa8p+B69OF3nihFmLPnDtypkh
8rUW05JPpUAlGIwkgIFwJ9G0deL5o30VcwK4cnDrgY39cJKmTZk26QCs7gXgTmCC
wn4F2ceXUDED5QK7Xv4lkLYCxvdNG/CtMq9tqYHaHwIwSTjU9HuEyjG1PiUWMJCK
YosvDQlRUXFinjOnDeZjHmdWw6fRJrOTUO51k56VkNn50c05YuFVatyzTbvL4uyZ
ub3Sm1jnO+FyMSjKc7h7xu9DQ/Jcyg586doNQuap2nnPWXlKfDrSJSdVVPpSX+2p
LXvY5K/QOYZXYg2oAIq7RABNXBrecBDII5PFOenPbNDVqWwGFQkX53xwuVPG8eMj
aRr7Kb3tTho9NiGGPVedDkW0zLqjUZNk+Nwm5uNBeay5P17tx6Po+pFli9RmW8cK
hlKfHrkgJtpF996nHZGE2BDRe1flz6AVVYXbpVz+xeedJZtySJcnEdGr1ksie00L
TbRehzxJyALJXnIZ9/BJXOVjnA1NVJseV6yuHXnwX+L3wIaW9P8h/h87pkIe4QUA
AA==
`   }
}

