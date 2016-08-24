//  ResourceData.go  -  Generated resource file.
//  Resources - E.B.Smith  -  Tue Aug 23 16:51:48 PDT 2016


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


func (resource *ResourceData) Com_blitzhere_blitzhere_lab2_development_key() *ResourceData {
    return &ResourceData{
    name: "Com_blitzhere_blitzhere_lab2_development_key",
    data:
`
H4sIAJThvFcCA21VR8+sCA688yu4o1UTG3iaXYnc5CaHG9CkJufw6/d7by5zmLqV
y7JLtiWzSQky6zrX6bbmCwD+oJjrvP+0l5F0+S9Q+BzJ/AFZ0OnqtfqT0A5Z0qr5
JfO/QBQFURF8EqAogCIMYk9Q4EHk+TuCEyD7ExRBXARF8reEIr8p+wRZGuR5EPip
8Y/mv8C/jOEf/H/Af36DFSTZAG2HAd+27DOuAKpC9EcBdFkWhlJmGUblGEtghqVD
P/FoN/tcI2jy+E5tETh78nHe5ZiiRWBHhwJF+uk8Y35cSYCpTS5RYL1efO2VT9kV
PyhC+9SwI9vfSL2ODjqXXXy+gjsz2aGBtxPTeCFXoRx3jD0G5nbkCWpKXC+pWGH3
Nk0sXl/48Glq46loqKc6zI7IETXc9F6W5tEvCZ8IzvVYnbCkCMAt4eX0Oc+Hj7ex
405jUW1XIcv3x9qr+Q7GkJqCcg+CHgUMNmWjJ49dyu0b64aPgKeBkHs3pBERm7Ej
GIFjWBFIoq/GrS2Vzr3KRs0H8Cmd2HLdymXSvt9h+ovLUoz2xlwyABQ56ge2vrvc
Uf2Al9wxmXENFUzSDHIUq2BY0J+WzDMWwzLDz7BZRSCOkj6kYozrPgOSy8vao38T
NgJFwjeuqjxPjKXKYPqE28DzMRV6YA4so/r9vSQ/KoQLad/DbazsMFwT8PK3Jw6v
FCKoocK8Uxi2momgMPaRvjPiqCbBn+mCh36WVTmEpqI0+sy+SFhPdkWrxASUlM3T
ocG8uYY+r8/THbJ6NMb1quBQRTcVk7K+VrapQW6tHySlkqBjfsB4xOSzzlQbAOl8
RlklhBYdfGavD6nxbP3aYSlEm5ioiW/K2WfV1qtEZ5CHxK7MnZndhlIO+2hEKcAa
9nLAuWPsR/D5IqucMtogtQ1p2CBC2t1qvLtZNc3Ra6aKgrlDP+v+hFInpLPC5Wgg
3cgE+8BcGQkMkvaC9IWz1YjdyosTZDcRnFOM7va/mDBZzd7cvQurnNjDEaZP9gW7
gKqIQb2ZcnzMgTdSMgPTUYeb99Z8v+WpJl0Fs0dVnLDycPeZgTdNOsLuRMh5bUrT
8YHjps9J5MS4f6OjEXmVKLd6VVvY8pzCezC0Vj/OJggPkTbJt3k7V4mGdvbH8bFf
A4CR6Ydvdk0/E2e62niPtHxs67jAD59yLX/o+q9m7Tg1BizZ5MioVxNNTmVVfyxz
Ui2AHc/kS+i+ffoLWg+vAx1coraX/DWRNjN9EWtRO4GaLHbVFG1T1ye2z/7hecOn
n2ETAvg6hnKI8T+Ypub59X64k9gj+wJHdu9k+M8hflSloH4sM4f+SNJJJBOx42ja
F5COfFhA2hzLt4Gj3HEFzx16HPeHfWm86731AkRg77VyIJOtTnu+q9BZ9J7JuiNp
FRbxuNjLAezE9jgVCie8U5erXzhqWqiwUmR464knnZdiZ/dqhzmLv2fvPt2jyxs6
tLb5JTu9ygCdW6/kLZMzYwgMRLoqgRbvQ2XLVxCQpHvdJC1FQngITFBwyDROpsrt
1ScqrcxrbYMC0goSCxq98SaBmebMHiHStWJ0656V6mhMGH251jIL4wKN+FiGPpeM
kCtEdVddxZ9cCcSHbZdp3BbINUw++wyKvPjckapqSo445ZrrO/JUjsewy1x55iw3
KI9k0GHVkbcomd+A+4JJZpCYKE6XeSRl2z/CVnjTrfjdirXQQnILn5i+wmQ5ND1B
mT7+6nrusR2RiRsGB2i9RsOi4aCpVevwwdafQV0h8gM9b0k8prHXo7kgx6xZdiM0
v8bSKzHF51sjSoOotl8gKljmsJpUPNWYCMw42QT1QKDWezKGe9skpvQ6uVd83Y/X
8kCXD8b89+9nIxj8v7+a/wMOzdn2KAcAAA==
`   }
}



func (resource *ResourceData) Com_blitzhere_blitzhere_lab2_development_pem() *ResourceData {
    return &ResourceData{
    name: "Com_blitzhere_blitzhere_lab2_development_pem",
    data:
`
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
x/s/YdA/iIsJAAA=
`   }
}



func (resource *ResourceData) Com_blitzhere_blitzhere_lab2_production_key() *ResourceData {
    return &ResourceData{
    name: "Com_blitzhere_blitzhere_lab2_production_key",
    data:
`
H4sIAJThvFcCA21Vxw6zag7d8xTs0Yjefo1Gon1A6D2wC6GGFnrC09/Mnc1dzNlZ
tuzjoyNbfNSwsG1Lm+9buULwD9XSlmPRf+3HUP6BleJ8LAUswsHQbs3fBf30fPRG
+dXlP7AMYFyCWRYWBFiWYYGGMQBzCszTsELAMgYDHOZwWGFhhoYJFpYUWAEwQcMc
gKFfj38M/wP/257+Ef8H+td/ISqqbsN+IMCur8dCqMCGkv6dgSxdV6ZTFwXBkARP
ETAp0PiECiypTaOQwxBQbP75CfkbQqKYpyCRCArsIaotyZNVYUJGPd1fyZHasxh2
Nh02ki9X642i2WqQyD3aO1CLNz1oRyU8kzKm75czW/29ouIrW6QWAlh1IgQIh6Yt
aU5JUqMz96DemBEd8unBvR71t9EudNDQKebEzKokZo5n+8hqNXXuFMShnwM9Ol40
FFXgk4PgndSz5eVTfw8aTYg29f3EwJ3HwGfdhexZzp02WBXBoWVq1z2oL8YNT1p/
M99SxFZu2b1LbRpErxTjpYl9aalDZRt3twX3fJtkdad197Snzs572WkzyEtaoGLn
C6/7WmsCJvxOdCaUY7W6TyMPKDJHpOTUZcETRGH6ie1Jp14Ml6/t7OwyJsS9DISU
vjNrqgRe9CxPAlWYecDkagtYjLDKktSTZ43hR/aq6Le2TbeqpccSSVJd9jNIvqQf
Bd9xivtjCwal6RLBPp+tkXLhQ27eT3pNRlAjnutWiDocC2I3uLrenPp1UjHNQlsQ
9OveP3FUm0vdkF64uPjK19tOfX9Po8qcswj8+3j4FHYhFILFtEWQfV2ZhEdjQQd9
zs+lqUtoYrKO1Vv+Katuiql8OThKVvSCuY4KdVZ9fJOm0bPEg2oq84m/IqIe6FZ9
Qh11WcWX24wXLvf5z4TZ6nRFiM3mcCs6rFEf5f4B7hkdVmsZ1bLGaM07I4YMgFv3
6oKubfkpSjXCpIqCazyYRRlp0S5f8osU3vfgc9+Z0rIRWSXRrzVWbh9+0Pex7CuY
bONVQ+959k4+YG66j7Ll+cxSbO2BYN7rfFsT31wI8ZoyLpuDI65wmW9nM7yvsrCD
ExsO34XSCOHulkOwVPdywvo8qK+wLWBampFvCGe+3dTrsPfjUbCidaJjkZkcb563
vxnbewvVlecT5+HZnnKXXDJpzbBKtCiSRy6djT3+LcjTDUvFz0h4ZrJbquCzxuT3
zrE5h4UnRFhJ8H5Eg+fRZ9ZU+B7crixM+KK1/QHYCPAtOWc4aUMSNdUFg3m2/BGe
0mQYbyvgbQjlnjQInnuVJ9+t65AZqUcsXp9jtt0krIhq/xlJ+0f+URZuNzclzPxC
rr0Bs1qlQX9B2qGKeS6U1cOipBxBWf+SQ+oq5TESJmp4FdhIhEZBiUL6PsUg5D+6
2o/s6UXdoi5YADUnr0slmTcWKI+ctu75O5jLtZvM9217p3Ho9F7+Jc6tsJld4LL6
9fWvw5UK6UEUKP+AtFbtou0pbb2XblzdCy7oOquT6lREv+vEmGbPUG2KC1I6Grx4
l4OLwF7tpizVx20NErLFMdYxpUt39QOUTLpAnYsp3ySuvH8FL8VchrpbBP7y2oeu
rE+giYsK7vrys0GGaxd0tnVcKtuxeN6JSm2wa446XLLMiwly04Sg2YSnIF7cQZlU
9zEMNEeRgnDXoUNuS4SjEBCzhPEMsQaWiD8WrK/8D7UodK9LsSIgVLngmsZgbEdW
Df4dUUaa/dKVo7mTdKdyADSROXnDge14+9t+MQXXFkusZVUkX9M3kCdbkcy5+ajR
gPfPAiVXkcm0U37FGVpZ+F2HgJaIrYOCOV+F390n+sptc/eNqB+/rFGh6VjgnBza
cttUj2JeFHfd+d+zUWz5/7+avwB44llTKAcAAA==
`   }
}



func (resource *ResourceData) Com_blitzhere_blitzhere_lab2_production_pem() *ResourceData {
    return &ResourceData{
    name: "Com_blitzhere_blitzhere_lab2_production_pem",
    data:
`
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
`   }
}



func (resource *ResourceData) Com_blitzhere_blitzhere_labs_development_key() *ResourceData {
    return &ResourceData{
    name: "Com_blitzhere_blitzhere_labs_development_key",
    data:
`
H4sIAJThvFcCA21Vxw6rVhDd8xXsUWKqgackEp1LMwZM29HBYDo28PVx3lski8xi
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
H4sIAJThvFcCA21Vx66sCA7d8xXsUTdFhqfRSORUUOS0I+cMVRRf3/d1L6YX45XD
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



func (resource *ResourceData) Com_blitzhere_blitzhere_development_key() *ResourceData {
    return &ResourceData{
    name: "Com_blitzhere_blitzhere_development_key",
    data:
`
H4sIAJThvFcCA21Vx86rZhDd8xTsrQiwMeUqivTRi+mdnammg+k8ff6bbLLI2R3N
aObM6GiGeZcwWJZvlaxLPkPwD4pvlfdZe+rvLv8F89n+/mYwAztdtXz+SWiH9N2q
+Slzv2CUhFEcxlGYomGBh+8kTPEwRcF3AgY0DFCYfcI4B/N3+MHBQIABgAEOC+Rv
Cv3U+E/zX/Cf+vAf/hf0x28wvCjrsO0A2LRlH7g8rPLRPxFIk2V+2GUGAJUFFg+m
asIM34lE6whKU0EtWqk2jJ+zyzH1ox9NKfuGyBadjigUnm5Bb+HG1/UcqFw6hyrW
o/PH+JlSID+p8SLvoPBjgxkTx/l49UzHlV0L/TJ5zzjGhEtMW8gsvZfo4DcOJcbp
s5sKRQSz3bBVsC17QTSe1Pnk8dQk3vDDnebHpJWc4UFOWTHLGPWAEin93ISxi113
GYaER2bWQJ15NEowmIk9RydR0HbIiGwiK77wpVwhdsvnvpMvLEZIHxpRkRZCy7Xd
1fG2CnM/q+FbZKOo1nN0sYRmqZnAoyTSrLXa6YYO3kQr9UJYHvdbFJPQNN+N9vU4
slUhBKqohqaig/waq6speUPFbiget7vMAQswYPhZNjMMRp8+GlLEkBS3oGxVb6Ax
36Hkr5mOKZnwZUQlFj/N6H3mkSrF+4Tay/VMZP7uvLKH0SGLqpkSwZLtUM2QNETt
zDiI3gJA0Q+EEcrnHRGWVExxm7Tbx9DNrDah4i6/goeH9lshFDSDWUnlsEqZQHGv
lmERCqvUhnu1U82YGM+aFcbveva7YGXP0usvrdV8J4kSUtnygtAV2gxuo7JeSgTR
u3iFBkpFzzMvEAPzhAdqp/XURgz5zHHjKy3akCNFbN3Wm3e8dfGZkeuxacFURlwT
QcHnGYF5ui1FU8SKmheOGGVWcptCabKfsXdfGpRML/qxTxkQ6RSlpBbZKvyFmSqR
7Q8Im++IM/BsGfEA82TB0xZkBqa+KFI+z9b6HWmueF5tryngYhCcOsqsc5KNotH+
3XwbqBAdgCotFrM+u5tEADJE/xaBR4RZrfiZWNnmJe9HkWYmRWlt4jXkOU6Z0kVC
az/kCxLXZZ6CEqFNlpv2h6FXmxWposlgL6DGD5lRWqTqPkH8VVvTH+d90uUz+Ffx
xQc09Kx9i8i+Tp8K7rp/RZog8GBf77T+xNdmRFnWXr05cN+8CzaVPnMP/Fjg+EQb
r94LhYPcOlkPyXjllHBUEkn1LENWGltmNjNhfsc4RSW3XOvERkmQflNqDWO/V8vy
1fLQBN+BhIxOfJVnY2G/CROPCkUTJLkY15dS9qaVK0sikEX6I5kxym7a63UcZPTp
9dOsHdvUQINGx5M6cvFknVaaChtR6g+DdKfAovxuo3P8rR2OHH+uoed1dTqQ2v44
U6S/VO/9Lh4Qa5ddD35MK5CS0/UvR9JumStSjnxHutnm5hRvZNpE7nG+9f2IqMZe
YVu2ERQAaHVc0GZ9c3XgqYyQv/UpfsS6NwZLZUqLjR6vLeftfsImDgULje9F/eG/
pi7UZVNJuvJwFwzKxZBdnPvT6mUBuEN7425nSa4fUMYIKr65rpN9Z8WLArWVnQ6K
hdeyhWWLIKvlb8QvEPnqdqfm4zezDdeNqcXecVkkilrAXm4SrlsT9B4jLa8hSsH5
mRRrzVFt0QzZVHmpSiA0p1Zr/5EsBRKh3hqk373xqHUCHMqp+nh9Jqk6LQ3JvM4R
pws6PJvq5uslSTThDjyIdRqzVOgjt8eOb67IHDvOeJeVF74z1cLve9qq4V1dM37B
Mdpy0+AmqjOPsC13nQZlQvn9y3GanbXau/+5pyEA4w2Tz0qkzlHk8M4sDVmZzrKI
kvZS/eNggPbvs+F17v9fzd+UNBWuKAcAAA==
`   }
}



func (resource *ResourceData) Com_blitzhere_blitzhere_development_pem() *ResourceData {
    return &ResourceData{
    name: "Com_blitzhere_blitzhere_development_pem",
    data:
`
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
yv75Vv8byXMas2sJAAA=
`   }
}



func (resource *ResourceData) Com_blitzhere_blitzhere_production_key() *ResourceData {
    return &ResourceData{
    name: "Com_blitzhere_blitzhere_production_key",
    data:
`
H4sIAJThvFcCA21Vx66shhLc8xXskUUOc2VbIg0w5DywI2dmyOHr3/H1xotXu1a1
ukulkopLKpBd17lJt7VYAPAH5dwUY95fRjIUv0AxP5I5BznQHZq1/r3Qf7KkV4tL
EX6BOAdyFEg/QFEARQQUeZBDwScGEk8QZ0GSAjkG5DgQQUHiAdIEiOHg8wlSBIiL
IPBz4z/Pf4F/Gp//zH8Df/wDTpQUA3RcFrQcJWA9EVTF6DcD6IoiflmFY1mVZ22R
xaCx8CSipmee3t4vdMJMb287lGw1Wr2OiI3IAbqLpGCxhNQZGKhGwmZl9JxZQ9hS
UX1L0xNtoYrGXKvqEc//SoJOufWeelNn2s9OiGsnuwTCudc3Bp1AdGkJxZfOQ+kR
qmk2tteOXdPLPZ/4gRHHOaylEtqkb8QI6jwoTDAs2LoxGZzoPTUHAJ9s/hHa0wzZ
svlCpZdu2JNDsWdp7l0U0k8csclPKo3tRF9l+EYLLQ2NqEEWQoQ71wSQIzHQKlRP
D6qprRBY1PvqL5oRM5L/hiZVw/f0gtQzIE3u6NqPbmH6ziQwzwUklRU9oOZz2UCW
FpTMODFkR0XBO+bGRtOWYcZV1mIgarYVgbVZjv38Y/bXFAVzmKJXh9JZDtCWOfCQ
oQm2KhfHZmGDXTmXHxUbplnJWKvigMnKHgcW+ipz+FrW5n35Jk1EEGaQjwAw79B2
K7coUVlIiBgZzE8ntp4rEfBK1WF93swCw7rMTk2NitAxdxTctXGQjY7PwzwEVMGJ
YyFeyopfNu8cEQUM5+zBTqReHc4qLsOm+Ljjqu+mXjhPbvwaJwdneqLKFIR0IpA9
kC4r+Q4TPXnzvpOY7m20n5f7Is+aJmsnnSYebmSZm8ZRsQNjrBV3lA8qYZjqPHbg
Fl8m0vRG/DaKuFqRtOXlU62Wx08sKkKXTEhYvLOvNtnYF+ZsP9nqXrrOCKtvnYqI
AcN9Vc0k8lUkskwzR89X22zfzN6S5YfydDhjbu6bSw5af6GSgYPj2TeN5gdBbtVm
qgB8lG7HdrvF98wO/0jsw6LE2CkEVFlUZSYrIwlOrb0ckVSz2lMQiT23jyXzarwb
6vQFpoRW5jMY6kKLs8Tbyg+3yZYwioJ9aGNrTCGpJtlyfGt4RR0qIFCHuQP/t2JK
HjJgccrg4UX6nI5RD0nuufT4VGkmjIcPN+o167kuAnuPt7m3F3blgRwih4z5y1a6
+qh2wGRrQa/f0vKSM4npHSpkc9zIV4Eppk+Kax7aNEQXXu2DZr+TgsKBdqJccDqS
EvObwwF6mBJtVm5wYnOmJNBR9lFLm6l7NSV4eyHPlPEecPcjmY3jl3ujgn0JLz+L
NhzDPogI6DzuNOdCp7C84Ca9hy+if5ZyTthy3qcBwncf8zFGdO+M4nUQAWXcb+qg
HVy2EPqJrYBl6HKdIvUo2xzT74wVj6SnVwnPDefzJ5LeEupB/6n3B/RQtgjWorrw
FWd6fugQUW0WkPjm+cZEvhkSg+qZXENSibBVrrJ52xhbvz4t3TFi/rN97UeaSRNL
JJ+675g8CobzfgNQqDbJqfuo58BNFbd+hE4xfX+p1RCcrjNUdD+1TJdxxpreOk7l
x7XHeZU+TXWlPOIFyA/bq6FMcZbigZ5XR0qvTfCXs6QC2cdy8rbfed5cyBjvrKY9
OzqAJ/ezzeHXfbQEuwD2VDjNv5Kp+qL6VT/arKEqeeh7lAvh6vxYqjNkZSCrphA+
OU+yYsfzscaDD01ZgcShVlK5X7meX8Ls0Cxfe3rCrkZuc9NhtgvSC1XPqBLtyPTm
8hfEsDk7rp5u64RvbgCu8LuK2/raY4ZOwOH1VMu6OtKDgaUgb2Ypv8L4ntKepaRo
7l0CmULo+OuvfwtHNIT/Xzf/A52W95EsBwAA
`   }
}



func (resource *ResourceData) Com_blitzhere_blitzhere_production_pem() *ResourceData {
    return &ResourceData{
    name: "Com_blitzhere_blitzhere_production_pem",
    data:
`
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
TUiTL+76Ur6/SASN//v3yL8BpyJ640gKAAA=
`   }
}



func (resource *ResourceData) Server_ca_cert_pem() *ResourceData {
    return &ResourceData{
    name: "Server_ca_cert_pem",
    data:
`
H4sIAJThvFcCA2VUubajOBDN+YrJffoYDNgQTCBAgNhs8dgzNoPZjBfWrx/cr5N+
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

