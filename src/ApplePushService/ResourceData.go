//  ResourceData.go  -  Generated resource file.
//  Resources - E.B.Smith  -  Thu Mar 10 22:43:58 PST 2016


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


func (resource *ResourceData) Io_beinghappy_beinghappy_d_development_key() *ResourceData {
    return &ResourceData{
    name: "Io_beinghappy_beinghappy_d_development_key",
    data:
`
H4sIAC5p4lYCA21Vx86raAzd8xTs0Qy9XY1Gonwk9BJKyI5eQkuAUJ5+cv/VXYxX
tnxkHR352GJSwcKyvJt0XYoZgr9RvptiyLvDSvriF9yMf6dFM1R1Mk3HH+lf+Q+4
G7Ok04tDlX/BGAkDFiZlmKBhioIpBRZZmJVhQYQlGiYoWJBgmfqNETCYYWGMhjEC
pghYATD0nfEHkV/wP9b4R/0v9NfvEMFFtWDvJsCOp4aCD2AdxD8dyFRVMG6qKAi6
JLhAGId3n+6EbMcelthtvszaJbVJTe+jzX3fuEdZP2mF3SbMM8FbXiAvUv1+uYbr
64n6rogj2I7n2zxkx20QG9ZzjO69ckpFakF/dVr1YlNeL0Stk1+Z0gpjqCX2N1Fu
yTal6i0jDSOTiEveSrY3Oj4ZLnR5Xr3+DJzX8ZFLXfPjTEnzsTPZ48NR/gLt4jUb
8LHJNz0LK+Nd1Yfs+jlzMKA35KYvl7CpKzYtg0PbTQ4EXeA9aROvQu3KzbsC4TTu
dO69ZFHXv5THxTH9Nddp7n7ax2a9nDnTA8YJhfqqg25J+ItFp/kNULNAdq+URKDY
4jDuFqOxbo8pLzme+xSpfClW3f6k8j68OSc4XFUWXEEUxq/YoHRXjV9DBzOd2LEg
Nx53eSBFAL6qohzJ3Qcw0wutnPmQPLzlHmhj3yC0+Liv7Tz7mmiIfIqzQXIByjyH
EHYp89lGHuJExdhgKDtp42lnJcGIhs27oXKeD/Ds7c59rpzioiMpNd8mOZPtG+J/
thjiAW4H0lj3jdaNB7UJCtVKi3cxe8VWskruI8ThU8x6I1RIfuRxR4xgbyIXoDx6
ic4Rukp9M7Hq4aTR7iYY/tIihmtfruX10t3fGZoAqh2+DKGtRcQ6HnfZy5rh0jbk
y9eor4jg/mj6+zIYCxG0V0I7CHlrmNCw+VS1utUv6OM5G3ZZpc3b6pEaPSZeIJ0o
MRC8j2oEwiIvxytMqmIgnGhwT14H+QKz7Is6j19p9tMUNu3i99UxeLM5MXAA1aRQ
1+JVQyjaDnpWs8H3gnQjkXcnGYIUbSjj6OCdZYmpLaJnLQEYmFMyR0pWkXWWngtD
H6fLdn4gUS2kessAzgd9jUIxreSr0GcPc34A/m6h+7GY2Y4aL+b5fIYe4WuT4nOy
aT65H8aH4BJQg77RJLH8BmSpaSoje+fc6rSKW+P3PmgREucDFH84OYWEzFa+B4xQ
DDwR+zIfyhsFoESuOot0GrG09XRd7UeqzRrQvU0zv9ZXq8mPuMyyVBfIISsJmtVY
XG6E0cgP617zCcSV9cjYyl25Dai+0B/UrfPWRo0NYPwTT65qBGI7yn4ot9b7M5Dr
I/u66SS2VyheHYj7MGZ4O48ExWm2E7cm3Qee9WW/+NSW+dnVswEuy6rykZuv5nPP
tgn3CHwmLXTR8LmGhkulH6i71+LteW8I7roOXT1xCO9owKzPVRikU+2m0JyHbb8X
YaiE5l7he4PkVn66PbTUa2eOutCDvA1iCzs2Wtm84EtZNOewzZDEQZersl4mYn8c
Bf90hISdtJZ5bGNRySoUkDtDkl4pvenYCV3ma2ypTFBG6F320rOYfs4ef9sqv2L2
58lMZYGzXsZ9XMf9LsStgPBLO/FVr6kPswMIErgM5rDbfmkjX5ydsbNYm3H4YLGw
cUMtW0T8allOaZNfZNPlsQB1scmHmy5W0vO6HbGMMrItFPw+gmymyqt6pqYfqUHy
2BrdQGTMkamjICgtpFTPyX0Duq+Gh0SJZJKPdC+51yhRQnBMs+TfbKDvAiuuzgsJ
FV7oRFAnLicEhAvmFx+8eIIXYuhZ2M/cy7Lu8PjpI+4472a4QnPoodX0JoRhYkt7
qruO1RnL1CG29z3lPy8FWPL/v5r/AKR/xmk0BwAA
`   }
}



func (resource *ResourceData) Io_beinghappy_beinghappy_d_development_pem() *ResourceData {
    return &ResourceData{
    name: "Io_beinghappy_beinghappy_d_development_pem",
    data:
`
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
6gK+X23B4P/5Zv8LCOV/w3wJAAA=
`   }
}



func (resource *ResourceData) Io_beinghappy_beinghappy_d_production_key() *ResourceData {
    return &ResourceData{
    name: "Io_beinghappy_beinghappy_d_production_key",
    data:
`
H4sIAC5p4lYCA21Vx86raAzd8xTs0Qy9XY1Gonwk9BJKyI5eQkuAUJ5+cv/VXYxX
tnxkHR352GJSwcKyvJt0XYoZgr9RvptiyLvDSvriF9yMf6dFM1R1Mk3HH+lf+Q+4
G7Ok04tDlX/BGAkDFiZlmKBhioIpBRZZmJVhQYQlGiYoWJBgmfqNETCYYWGMhjEC
pghYATD0nfEHkV/wP9b4R/0v9NfvEMFFtWDvJsCOp4aCD2AdxD8dyFRVMG6qKAi6
JLhAGId3n+6EbMcelthtvszaJbVJTe+jzX3fuEdZP2mF3SbMM8FbXiAvUv1+uYbr
64n6rogj2I7n2zxkx20QG9ZzjO69ckpFakF/dVr1YlNeL0Stk1+Z0gpjqCX2N1Fu
yTal6i0jDSOTiEveSrY3Oj4ZLnR5Xr3+DJzX8ZFLXfPjTEnzsTPZ48NR/gLt4jUb
8LHJNz0LK+Nd1Yfs+jlzMKA35KYvl7CpKzYtg0PbTQ4EXeA9aROvQu3KzbsC4TTu
dO69ZFHXv5THxTH9Nddp7n7ax2a9nDnTA8YJhfqqg25J+ItFp/kNULNAdq+URKDY
4jDuFqOxbo8pLzme+xSpfClW3f6k8j68OSc4XFUWXEEUxq/YoHRXjV9DBzOd2LEg
Nx53eSBFAL6qohzJ3Qcw0wutnPmQPLzlHmhj3yC0+Liv7Tz7mmiIfIqzQXIByjyH
EHYp89lGHuJExdhgKDtp42lnJcGIhs27oXKeD/Ds7c59rpzioiMpNd8mOZPtG+J/
thjiAW4H0lj3jdaNB7UJCtVKi3cxe8VWskruI8ThU8x6I1RIfuRxR4xgbyIXoDx6
ic4Rukp9M7Hq4aTR7iYY/tIihmtfruX10t3fGZoAqh2+DKGtRcQ6HnfZy5rh0jbk
y9eor4jg/mj6+zIYCxG0V0I7CHlrmNCw+VS1utUv6OM5G3ZZpc3b6pEaPSZeIJ0o
MRC8j2oEwiIvxytMqmIgnGhwT14H+QKz7Is6j19p9tMUNu3i99UxeLM5MXAA1aRQ
1+JVQyjaDnpWs8H3gnQjkXcnGYIUbSjj6OCdZYmpLaJnLQEYmFMyR0pWkXWWngtD
H6fLdn4gUS2kessAzgd9jUIxreSr0GcPc34A/m6h+7GY2Y4aL+b5fIYe4WuT4nOy
aT65H8aH4BJQg77RJLH8BmSpaSoje+fc6rSKW+P3PmgREucDFH84OYWEzFa+B4xQ
DDwR+zIfyhsFoESuOot0GrG09XRd7UeqzRrQvU0zv9ZXq8mPuMyyVBfIISsJmtVY
XG6E0cgP617zCcSV9cjYyl25Dai+0B/UrfPWRo0NYPwTT65qBGI7yn4ot9b7M5Dr
I/u66SS2VyheHYj7MGZ4O48ExWm2E7cm3Qee9WW/+NSW+dnVswEuy6rykZuv5nPP
tgn3CHwmLXTR8LmGhkulH6i71+LteW8I7roOXT1xCO9owKzPVRikU+2m0JyHbb8X
YaiE5l7he4PkVn66PbTUa2eOutCDvA1iCzs2Wtm84EtZNOewzZDEQZersl4mYn8c
Bf90hISdtJZ5bGNRySoUkDtDkl4pvenYCV3ma2ypTFBG6F320rOYfs4ef9sqv2L2
58lMZYGzXsZ9XMf9LsStgPBLO/FVr6kPswMIErgM5rDbfmkjX5ydsbNYm3H4YLGw
cUMtW0T8allOaZNfZNPlsQB1scmHmy5W0vO6HbGMMrItFPw+gmymyqt6pqYfqUHy
2BrdQGTMkamjICgtpFTPyX0Duq+Gh0SJZJKPdC+51yhRQnBMs+TfbKDvAiuuzgsJ
FV7oRFAnLicEhAvmFx+8eIIXYuhZ2M/cy7Lu8PjpI+4472a4QnPoodX0JoRhYkt7
qruO1RnL1CG29z3lPy8FWPL/v5r/AKR/xmk0BwAA
`   }
}



func (resource *ResourceData) Io_beinghappy_beinghappy_d_production_pem() *ResourceData {
    return &ResourceData{
    name: "Io_beinghappy_beinghappy_d_production_pem",
    data:
`
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
`   }
}



func (resource *ResourceData) Io_beinghappy_beinghappy_development_key() *ResourceData {
    return &ResourceData{
    name: "Io_beinghappy_beinghappy_development_key",
    data:
`
H4sIAC5p4lYCA21VRw+jZhC98yu4WwnNgFklkT56MZgO5oYB03sz/Pp4N5c9ZE5T
9GaeRk8zbJzDYFmm8rUu2QzBX3tPZdalzWHEbfYDLvs/X1nZ5UU8DMfvbpptWdMP
bdYtv2BNn8SNlh0K/wOmAXy9wYwI0yiMiTCPwcIN5imYpmGMh+krjAowz8HoFSYw
mBJhjIbpb0b4CYS+PX6j9AP+y+h/i/+B/vhprCApBmw7ADZtxQeuAGvC81cF0hVF
GIDCAqBxwBIAHjpkLo6Dg3Fheo9RvF7HLSSohhbHpdVK3ZOLEmD1q8YT6TbSUIDr
83w8WEmIVuQtcub6cCwv6Ss25Hnyo5da6aluWzaT4Ljh6KCXWNzvMR4JxazxAYB6
mtxjEWvXZypFyHa1XYIeMZkfXs5NfqwForlqHoW42/rJjSJyYtbliG26q+uoahyw
UNdfjFSNhsSg+E7VCSsoN08iY6EKVW72OWs6rw/NRhz/ruI7Kly1ZX6pyIOZPoo9
TBTEvrbpCdbXWieLzfT+Ucve0xtXreWzmSoaSrpiF03QDZs1w6xvMa7qDm/ErfMc
LBJvoWXramE0/Z7hbGcWJIHUu7aur5pyOoiI1vZZ04ul8MACLOi/y7Z4cEjZnnCI
4nUSD5XdRHp3wrtN2xO83ct9FFvhkrgcYq9lHBhO7hXMKg4chrRtf1dJd6BmdScQ
+95HDVVAIjo/WgWXxG2Nh5UmLrh56M2CfM49whP6o8+BTXraM6Hxd4AtHDqpCCdj
eEthCZ4kG3QoGlU96/vS7EhdGxfUTV5DLUl32vSLkG1qAVvmIt4bSlnv8XmfXyyD
XJdQFnXNBF4EjdQmIPP1mWtyva3LR1v48G0+euScgqdOJMuuZXmnuLuUpRt6mKJT
ydlcbiUxh+C62lB1p31Z3rxd5D9qUMnUKYCuoJmwvt7YoVj0pXWrwBhmW+K4zyUe
bPZovFYDKT9evlKAtFfAnBbCgF5igcmLFzJGcHKL91LouPYzynqVObHZOVlLBcvF
sFGDuATpxK9lAALUgDAan7uHVGOyERTKor1eVG6L3r3Vp9a9HfHSpmzgyGqF9yan
yoiXyP1S97JwxXZ8Ag1km74755xe5AePLHFW9VtH7PHAOwZnyxpq4697kruaSllD
eqMysnHWOZbdX4wfCQaZlUCgE608hDd4vY7HwJrDGXFjn7RnYfNb1D3swfYXLrt/
9wjiItgUibcuTZi1BW5FkJweEx6/UfHi2k9/saUsjue4jz66oVaqcqG71YwGEN36
tDbPQ7o1WhvehvSzU6vv8ApEM488eKy34k7b4yWvvMGvzRQVzty336R5uswr8DUm
/FIGssck/bG45KD36GSHTEfNUNQwJCDcMlEWklNk4Nqf+6f0+kdoRb6YMHyBELpd
qRMoKSPTui642UonXJm2JzgtFSDUqdVgxJ97EuHrcKu1mXdzsIKM8SstVZBBD5j8
moV7Nsjt0zX3py2V7E0xJCf+KGYBLd5bIwkyySyf/oos+vCifOk8Ln8KYCUR54mz
1ynPHoRsCHihTN9bUrpdHI2VSfsmUUDpLsmqmXDP9Op9h6EmT5BdYxIP7zhfAmf3
pUsI7poa/GYvWg6OWxM6iYjVNRc7zwyHrJIvjqM83akTdGx2beVcOOeUNnKfnGWs
Eam0FMnbW863nDEisk1RiDSc/CKTCbQB0Jm/teCYv5RBrcTq42XUdHsaxAak0SHZ
j3pjDe5o+socl1W/5EDxraSyiDZq9VxBFmg+5SVudDK8lK7/UrV8u6vuItlvaolk
DcTaqYZSzK2nPF9J6rgJlXfKY0d2wXC8W7BBZkrly+Nudd9jwGKC+2Tfqah+lurC
ABCPryiqRZOgI7nMiCzWUn64eU3+99//PRzB4P//3fwL30VgsUIHAAA=
`   }
}



func (resource *ResourceData) Io_beinghappy_beinghappy_development_pem() *ResourceData {
    return &ResourceData{
    name: "Io_beinghappy_beinghappy_development_pem",
    data:
`
H4sIAC5p4lYCA6VWybKjOBbd8xXsX3TZ4JGMeAuJUdjAYzLDjskCIwweMMZfX8Kv
Mzsrqrs2TQSBdHV1J51zBUwwC+73a5X29+LGsPQ5XqvinJPRTJriBwu6jhSsVDwK
0nZNcb6zyHLZr/5Wsm5xfVRZcfvBVu0faVGdcZl03fjb8G2PtFlCdsWIpB/sBrDL
LSso7GbOcgorcay8ZaU1u9mwnMRuluxcZiWRnS/ZBceuFZbbsBsqkaeNzK1PT0V2
/5z5SPr8Xy5novn5f8U8s/zPrc5vFZV3RGEmfvouU91ufXH9fE9m1r/to3P2x6T8
PQvaK8mHKv/lt7iyTkGSe9Web/8J6h/VWLG43qtjlb2nLOjvZXut7iPzr+mBsopM
VpQdDylIBJ78ljIGQgrxRBFqPAYDggAjZDQGuvRY+ZL8QbIjfdfGqHxkJrBlBdpg
wHH0FF9Ah9g8MBBEHiAHz3CMQbYj6WDbOwl0MNPgLXahnjZGb7jDsMPvtb0kmj/X
wnSh35hYy7tY9bGjHvg4eD4y9TAiRSepqsyTQOizlyIbMlIB58tgGDZ2CIeU6h94
YUxVZ8EkgUOQ7JC8OdzSBSRZBV1qqMxV8kjPBrb5w0jHTRKY37IKwjx0Wuq9y7V6
YLQyMw3PfxoS4o2T8Zq+wVsWTTLul+wkQiLJrSFGqogbV28bb1s9a4YWxlCp5X2j
kyRYnRNVGTKt7qPm0KVN3kY05MJTVEO2f6Yh/ExjSpv5zpsj6dnGri/QI4JcxmPs
T6E35BSHxhr9gwPm24PsGfSYJw8iLbglt6Mlx5pBdMny5L0B6vcaLA3xcDAGjOXK
AHNVdC8q46J0IdkyBLYPwBJBaQCTwg60FBO2FOf7BgQXsvf1fH6/eom7vgjH4yW8
BRc+uqAnYxcqkgKn9+IhHa/3NBrHF1rKkdNshK2PPpYrHTpjZSiZ8XWWR3TRbWIl
u+Dm6GbR7T1m1Ybgfo1rVx1aMxJPjzNswyBZlfP4/DErla9jf/Uzi9z3Q7Zea9vO
aCA5xlZ8eL2uhVQww5DXanAsTVSTLlKL05Kbk4YMi9bK6lcOk51gp6R1zazePX1d
fHzhZacJun87p5LtYebCP2/n7YZ/yKXaPFql/7hjcdWvnZHT+DLc2blTupflJTLX
r4satukpXLb2wMmykK+r1Gd21lF3r7H0ZRHpnCT8R2G7rRCc5rqyG+zIq8cicT3c
nsz9NnQdzfJ6DmADAqCeMNbIxEa49EA+MUuzl7KCbb+oq2MGZ87ajBdaaFsLYD/T
Z/6xOLeDOBHOmXsQoAFI4Mi8N7qGrEogwNBF+jPTL2kURehxW2+uvMOdFZ+4p/tL
FIG9nfCQI+DScTQdt0gj0GYQ15eyrlSBN2toywM+0jczVFG8qRQfChyQaCCVl5D6
8inpkkDpI97HKb/ETK6WXTZCClmd0jfuIl6ZU65XhQvL9FzjTFVGyjocheYrDzhC
dcuIN0mmOWXamIRBqtAgzWnppnmiHno0YV19ToaqiTDZYlJ0yqyxcU57RRoaOApW
NVLNB5V3zE/ep3yE89Ake/Fv0QwpTw0u6p8bSXZ2upgSLQqpgSmPTNPLaOF0U16T
R+rpHgereTYuh50dod0QQWj7GsByoAXlPNfAej8Ki3yR9W820lBpUYT7nn+HTyJe
fhjy/F1xbXCmXjDYFwhaaa1+qb8M1HEYEyZVhSEOUf9XQ38NM+VX5W9t7JEv8jpr
zHLfmCOTSmD/DaIt5bQhwiUY5OENFgJp+xHxbykMaJBBtPtu8/kpBpFEW5otTr3+
v/R/iQIFAjUS9cexCC6PkIuWJ4or4TiLvtZjZ94M78GI1d7qFVsS03scbkT3ICgH
/MDiruXUusp3c40IaHihbjMPTgvz64maD8HKZpII8ytJB8aHY6x/Sf3Kce1kdXwk
N1UoXTSXMne75McOzW5J0TbDgiTm3kmHD+tSO8YG1TW/95x+xeCN+TrR/nTEXbJ4
fuDrjPS93nomumWJa3VdE57AWrS4K1wss2ZpPbnBVCvSyItkKXqAMQ9+haSB17p5
JzzSEVDayytJUcjHWGOvPgX0X0ItD5dTho6Fg+2I3kmuTc8a7GabmcDkCojh3JR7
bn2ISZtL/iPZHl204ORZAtNOAE8UUiLJlraps467XLv7frX8/L6xZVP6+339J1QT
fmtyCQAA
`   }
}



func (resource *ResourceData) Io_beinghappy_beinghappy_production_key() *ResourceData {
    return &ResourceData{
    name: "Io_beinghappy_beinghappy_production_key",
    data:
`
H4sIAC5p4lYCA21Vxw7D2A286yt0NxL1tggCPPVi9e6b1SWrWdXy18e7e9lDeCJB
kBgOOBj+WcNg25Y227dyheBfVEtbjkV/Wc+h/ANup39nZTvWzXOer3+m8zIVe761
0/jXVD/lz94oL038A8ZxWEJhjINlAIsijJMwz8ACDovMnyXgYCDCEgYTACYBjJEw
JcEiDRMUzGMw9NvxD0R/wP+xpn/U/4X+9WfwkqJZsOcD2PG0CAQSbEjpXx3I1DRp
OjUeAEMArgSmAEd8Q/mM0aQorBMV/uxKlCB8zV2v2MoPKMeqL5mnMYl/vsQb9GZp
M9B6D+XGAVFPKSxAgu78x04+CRJ2NvdV0rp5oyvxkjbd3eVpkK2FIq1IyfXK9aG2
Mm+8//YYKZKzNkfvsUZ+EIldLLn4zKc9km7fkOfb6gBlvDtDkpXvyA2l87nL34VQ
oMeHF8BL5tJ6bl+ETE3+tP8uqp71eAtCjc3EJrmWnI6URszvoTInvfZk1mXutanS
UQcCmbqDI+SkGaGIK2vp9rMuLJm2fbLExnFT6tofDRFFE8uZQUo8W5RaYmR+KbLS
Tm4KxdVIhg/0WxDpcBeHYt2cKaS8CEWLvuZY+RPhM31qInABD6Yf2YKFhFrvMnTk
DvJthfI46bnNTNDEYd/f63p4dPIl3jRNu7fZRzO7SmrunK7aRW8HVdy8OlypTTVE
2kTt+2JDtqExouMSs93aIpI3sej2t2EsGQs5f3SqJYFlPg2q9tOGHAnCd3xKCeNI
jGgcmyefUF6UaIrQSGrtvdEJuCSkVT9Ve2O7fauMRte93e+VeMdLZSz77lXri+6C
PcxTnSR3dYR2odRi/w2elmD7D6GlibTDxd2TsWcxU26fkAfz7rYvluO2aO3AjO+j
xbvu6eHD6hYJlAgBIpe3cAhCEGRqQQ90zJPjiouJi3+UE8e+kkQ+pNsj69LYlhRj
KO+fwNylZGvjxwJN4y6MP+B1KoELjRPGjdYkfF/9FDaiRyN16du63hW6H+TS8FU3
yZM4bV8qyzY+orRCN9UPmDvKeZw2dFscyT8COHzeqIv9DINIt5nspABn8AQxIi7p
+VGgjWYJXx9tRFvrhCog210QiPUtsVpXFdLw+9RCfBeurnD84NqzCLFIju2JgyC2
NUQpUrIO9G/EHymEUgaMQXtf9yb6FsxPckqm1oZX0xhpPsrnNuj8W9hiy/vSDJcv
CcqF3qKPGNCdm/WgPajJD/XavXxfit0UdSSw7bt8ozQFdIod6NuJr0SWcXrVa/u6
2kwFYg8RvnfiJB5VLvsQ1hzBI2W1+DiImZwuIz9QOzefHkk92/Ir1m0VjG7+g8xX
qYWJ07G4BMG0C8NsT96MIWAWVr+20VOUCAUo5+I7M3ZtmMIpFotK0Uz/VKTrY82i
PDpUmwYyGW3saEibAlv7FFJUJTKTS7GWPOPy+c719y937w4uVkfWz4NBw1p9vHdi
TpbB65TxvEWdDEdephavT8OGvqdq5T4X0R5lJIqzh6yjNa7B1zJOo2NDydrWJkTK
cpo076lQ4eLvCZIYr2MKGSP0hGxcm1CEPQcGV+TgS+ytwehJ5OKt7R1kYsbMeKIv
D6UHrD4fsUcXyEEOpjzkvYZRIQdRZaJ5xMRrrri3VN+E1r0yR9NtiLD0zQq5A93v
HZqWVkfLgOnjRxLnr+4hdVjWPN4mpJAZByaFBzoTftL4kauU3tc5Fl2rsXlec9tu
EQm+iVfio4GYr3xzvca5/ARjXxWr8tCLNHOezTbugzx/JpSV63MiNfO2t4hlb0GY
vh+ZR6UCkbRd1svORmhtaTw+YWu8LbzuoQPBro4J/bv2HEx2Ml9SJzjn0d89RZXq
0Ay92fd1oSRc/a2am6uW4/632UiW+P+t5n+YuTMnPQcAAA==
`   }
}



func (resource *ResourceData) Io_beinghappy_beinghappy_production_pem() *ResourceData {
    return &ResourceData{
    name: "Io_beinghappy_beinghappy_production_pem",
    data:
`
H4sIAC5p4lYCA6VV27KqSBJ95yt43zEtF3HjidgPVdxRUO7CGzcLkJuCoHx9l9vT
Z05Px8zLEGFYlZlkrcpcuYAxIsE43srkPuYDQeLnfCvzNqufZtzkP0jQ93VOHm9d
dk/HsmtJ7eCQx/tQkE5+m8o0H36QZfdHkpctKuK+f/62/E5Xd2lc7/KnJv4gGYaU
KJLekjIgRZFk1iT8JAWGFD9fW7AlgUhKNMkCcg1Iek1yEiluSJYjIU0Swz2p8nT8
Wnma+PXfjlwJ5tf/A3l18L54neFlhbGF7Ur48hyiHIZ7fvv63qwOP9NrbfrHK/i9
C7pbnc1llpNiPuV11+c30s7r+HX88G9M/zOMFPLbWJ7L9HtLgvtYdLdyfBL/ej1Q
UjSTFCTb1WRNAK70bSUMTZNrURCg+kRg1iBAmnZqFdY7H7piA2bRCvVdF2nFlJrA
kmRogRlF4UNYgA6R6RMQhC6ofdewjVmyQtG3rJ0IepiqcIgcqCeNcTeced6hb99e
FMy/fKeE1QciUrM+UjxkKz4TBY8pVfynJut1oshUHGzv6SJLhqQpgPYkMM+f1gnO
CY73me0zUWyWiAO71iS7zhp/SFhYpyV0cKIiU+opaQ1kMf4Tr5s4MN+2EsLsZHf4
9D5TLzOhFqlpuN7DEDXGqIzFqCQq+LaFLxv9y1YJ8LJInSGEioAaR+8aly8fFwIX
xlBw5n2j13HAtbEiz6l6uYeN3ydN1oUYcu7KsoHDfl6D/+sanqpPuAZ+9ROa7tIG
NvpLLEA3OukMhl2nS4figP97Qocro6C+R0xRELii3G8d2IugWzuVtrYv6dO7GJxh
DbPw9inSrPveIgjA0X61l/jZX4j7KyIkHQHmBLA6Aa8h2EnbiSqLIaJjWEwiffKu
HjxYaFj25WmlUpcPl1hCo7Dy0QZBD1bX6VBJT93LjhE3PXgoq5a83aPUrXl/RWvL
lD4aTYvHvcnBvePti1gnCjf+yEXfboUTJXXtckTe9ZI/ZDtYt6ZAlwfj/DjuFvuU
xqnYf0h6mBu7KpwPzrV7FtAmBm4VtRXvPNLN9tGkloV0eA60XdibjzyWrvtBACem
01cXWjgq89rnr+yhtjuLFWh49TnCUe6fu03k7NRnJs7K4xPt6aO9O39kfLO+d0Oy
Wx0Urfc3dHlbFZ0m9U+0NeXF3ViKyXQjcYjpaRPBwu6ukiK33IesZAabMXo5c22y
dzfywfMzU43C4xmmfnbNMaRZAiA+CCD3ZgIhtTAg9SJJJiIrwO0o17DdsI7DzuM5
u8Zscbh03mkM6cF9T6BqGxKoADAgrxCvF7XZCg0YA1lDrTw3/VgUqPxYbvcpUoPU
d/zdgcVjD8UZ88GmEMRrWL3arWEEZ14RnKviaAkbMpYMXAHy+DfPobabQwzIU7FG
zCgJDZSYr0Hrw4Cr8Di/2EsRsVIvmoIp29p91NRVeMLTqeicpsj33IFzeNIp/F+k
rEkngY9j5Spk/DnD/pDxEJEwIcqUotZUu4uCNQq/BwUnCvSXaCyvwEiRn5EAKTwf
Y/qERdJYKGS22F5jBC/ReMIpKiGdMt7wDzQqhooTvlC8X/y7PhD4HnPayK+JrH6e
SEUBXSetvewrrG0I3SB6aWEqADvMYtWmUrGb9mzGZk/uexoxVK5KGGp6w/er0OFn
1/qu+BlK4WyLUOgkAR27kP+VgHmrGPGWMa74W6L/lLFGpnGdprSpcUG3mGb2M2Sk
e/gSVQMMbxLNFlYbTUUicL/J4niSKILdb1cQMXksiLprcSmVLWNeIDIELOuaDERg
QnR5O2YKAsuTAcBUxTkFV2faFFYsJzRsJwuh5W4u+/4gPgE/ZiNFuGhkJRPwfecc
8qZSuWhjOqs2jERFANNpkx/4wZh71Xr2gOs/Qm3tSecJVMvgbSaz3xGUaRqfokPz
m7GnBcVgWycqtKkemo7h7/nYQX/zWbG8sm2aB1w99ltGF51m6PI27Zo7T0y+u3U+
lwu74Wh+FrZyvR1arbnIejUMj+a+ohJVWT8WKhDxTJiFZ2RRmlKoAiVPhWVICCgT
9kt2AonS8H0Ne1mOjwceM2ZbXpcgPUjWtt6IjmrbgWi7DA2YW5LsnDIZuVB+bgjk
KPt2ox4oNMmZH0iNS58SsEHGMXCkoppAmGAtWJXVh3uiro5Jmaz7mL++3l9syRT/
+b3+ExqztfZwCQAA
`   }
}



func (resource *ResourceData) Io_beinghappy_happypulse_d_development_key() *ResourceData {
    return &ResourceData{
    name: "Io_beinghappy_happypulse_d_development_key",
    data:
`
H4sIAC9p4lYCA21VR8+sCA688yu4o3lNDk+rlUhNajJNAzdyhiaHX7/fvLnMYX2w
VLJslcqWi4tLkF3XuU62NV8A8CeKuc6HrLuMuM9/g/X4K8nroazi7/f69Sd/t27J
/8p+fbel+pXle96N3z4f1j/d3ZjGnZZfivAbhFmQg0GEAGkcRPm/oYiAHA1yT1Dg
QIICeQxEWBBlQF4AUQEkeBDhQJoHRQr4GfEvYr/B/xjjv/B/gb/+Dk6UFAN0XBa0
HMVnPRHUxPBPBdAVRRwPhWNZjWdtkd20B2/oyIyKu0xgoWAIKySwZSbKR9Ylbdca
N+6EAWkizEruSwhgm4fbe7ob12Ew/HSW5uxSKh9LUBUUrGr2S6M4mNJ9L6J5wWHz
lUXdZ7j6leUwE28FMEoXahsCX9brInQQqfBIo/VKOdLbbSwwXWyrXCYzDtMv+aSR
12j3e/SEI3VYRYdCSSDl70CdKQ6ZzW+8ivjc4GQktFVJEjExb24ob2cqhT43HJ7i
ZNIOh7pQuaEeQigGtzEglsjbUumrcD/Lh58e1kmVhGnbpHOslet3DqpkkSmwixFc
Q825r/iNFJtdmJOKvi0KA4r5a+wprjVPPJ9FClbF7xKyAs+YGS3Frs8yEZXaisDa
LMeOP2ILIWJLuaw/nPeUOzxgwvGzoAPWhfL8HNnHRc/r25Qx1WUO+aLXRFpLNiQl
55SkXpf7Z7W1sHv0me83kzkQgI6Kt5EE9XmrEHEGnSy1HvKBSmbgJIPTsUMjz624
WzfQ7xZ/vEu3DSc/voOpRQRLhAFCEvSoub6Xkj7e2mm3SQidIffCr2HG1NH9nKkT
/yw+28ZRvB4GZTivtbmw8kd7FGUnoNWjXawDscAeX36KN+fbDiOXqeZ8wKE8C5wX
EPlzm15LOZ9vZd89hMyFmaS5XAj5UAPG+DO2Lf4VFkYNGOYesWRi5HwQYcStNS9N
eXfT2VZJni2xNVHCnWdmOCeOqqUnzvQNnP7DLTiRL0ORJayeTQJxT26EvJ8ENSLP
PajdZ6eseO26quE9jvX9qto1ii3Heu+FHgN1q86QHx+E9wyHBWthmy79ls/cBOHn
Kl4C/Zgi94ZDJuU8zCu9/Xk+ea/Bc8yyK8ECjEfVJ3yqoGmcECmx/FzEfcNjwqCZ
CQfpaUmM0G6b5eIp45ff90NoMruj/zC+h+MDBJeIBs71HOe22GCx5rQ1aZhFHM1b
5SNoP3y8KJmUPBG/4zJUY8UXIn00V3XnuQ+XAJAMtaUP4nrSiVf7NsasrdX0D2KH
ZEbbBsRKekHoec8qOM9LZ2w3FvVpS2uKkXcxNDZQ6TfCPEgvMsKPOe+Q2yn8EyeC
6hTDtirgj/4QDf4fyutxFtxJ0YzeqnFfQDI8XABRF7RdnlnNfnx3c2M8RTKERN7E
SFHQg2gX74UnX0d2BhVuGogy7FCboG5GFKY8z00FZDzCrWIgnSJ+v6I2Qm01qaHU
IUYrX/XxFYyr+awKqEWmxFjtIOq16RwdO7o5Oe5gGXhiIpUIPNyFHitz/Cflzy5L
/r4LVmJ9ObP6Uvl4IbfjrrkRxz192U41c6H5MrXWeizgmgSzOhNU+at49XEjwAxS
4d+06e7HwWPOZV3TYCVHUQVvbGInnlI0RTYd9oBeaBGqAIX7vEk3et7NDtYxmqxf
C5z08yz0o1Zkn+qTEaS50vnbx03m53XkaTWLbBLjcpZkzQxEIgMVtsaVPOR0ntAP
szhLeSQTpfXzSIPk89iYxh/yrLuJQMgY6nJKBxoCZQ8pMkY1QO1MuRitW7A25NAS
JsIHJ3iPixFVYkASQ+zifi1D2QdR01jR70+r0NplkFYXcDdfC8Do93WzLXBfGWNN
rMQafFRlV3rpPZRukEEUkQrMPrXvwj+lD/4Qylv+x2xEQ/j/VvM/PBhTNkQHAAA=
`   }
}



func (resource *ResourceData) Io_beinghappy_happypulse_d_development_pem() *ResourceData {
    return &ResourceData{
    name: "Io_beinghappy_happypulse_d_development_pem",
    data:
`
H4sIAC9p4lYCA6VVy5KrOBLd8xXsK/qWwcaGG1ELSbxtoADz3GHAvA3Fw7L99S27
5t7umY6ZzRABklKpzFRmngNMChrM81idljmfKJo857HKL1l7N5Mu/0mDYWhzWsyv
edsPXX6Zac1y6c9lKmk3H69Vmk8/6ar/ccqrS1Emw3D/8foOSzvlf2Qvi22fJu0+
v2viT3oFaLiiGY7mNzSLnkuJoSFPQ5kWIc3taLSmGUCzAo1EmhVpDtEMpHlESzua
mpZTnafzx7uniR//3ek7Mj/+z7jfLe+D11leVlgHCe/ow3OpapqWfPx4Ld6tf3nQ
LumPp/L3KujHNsNV9ttzPtJO3iZz1V+mv8L6n2o0yse5Olfpa0mDZS77sZrv1B/P
B0qKZtJIco6arCFwlF5SytA0uTsiBDWlAFiDoNC0o/VI1/gzXvcFFu1I3/exVl5T
E9iSDG2Aizi6oQfQYWH6FATREbT+0XAMLNmR6Nv2XgQDTFU4xS7UT52xGC7G++K1
dxCR+WsvPK31iYrVbIgVr3AUn42D2zVV/Lsm6+1JkVdJICzpQ5YMSVMA40kA450d
Qnwi+j4r3E+Ks6aSwGk1yWmzzp9Oa9imFXSJoTJT2uvpYhQ269/JvEsC81tWQZiF
Tk+8D5naYEotU9M4etg6SoxxbBhDBDh4yaKnbPVbViPYidKXgQoFFZ2r992Rr24N
RRJjKF17PXR6mwTcJVFknKrN8j22OAtuj9hlGsMrsPOdB1EUhV95kKg4jMl9BXwK
/CVDUD8yRuGp/iNB8BiHOktCb9NHXyQBv0SdP5y6rI9IHnKXe41UHkLmpJrtIbBv
0gM4z9JAMBmoLvdGXSrOQ3ftR/O7bL+rpmlQq4FJwaL5KpuKxLCCzzoDYCFg8+Cp
gIo9mUtg2ZNONJiRla4qt45EU5zfRFBkkoozqj01bWM+Nk4Ubi1GmLfXKVovx419
Ta/mHZsC+roV1ujudJQob2V4BrrVTbXmrLV2uHM1dVhF9aBKhi/A6pDlKyFZzr1y
Z21TREU1T2L7ttUQU+87rej55WFOK/68zGpxGjcr/qBSN5459HZ3jeVVrF9mydmx
2xQ9Qn3cQWa0hmSWNmO92cZiUxZbLuHGxY3U5ZYqkQ8v+EhpTqZcV5Ehlm5kRG/s
etUkUsF4nzp/P7vBFKCv98/bruAs2946eC5dv3VYLYstEUxmeKcupPMOicecF/ts
fems97lbn8fBvKabfS1v8lHarXRpmCIgIsHKeCVxfSDEu9TWRGADSPUbDXIEjSCX
sPoEk7OyIIxIOb6wHNkTWpTNVZhDfOuHh8ukyuZxMUDzxEYmYRtRBgBYxa+DNYQF
lnvgVSgNUdKxioIOO2vZCk0W33zd63cpLgrp89kPqgskMleefAAB6fN3/Av3WR0D
DxqaeiOvCItihMWTB1L01JuLO5yYJ1rb8tSZLaUpwqKpTp+ERhH9G/K8Iro0RRRw
nKbCMr04nKbIhGn8OX7qBmYdh3BFEYWagOIaV3CVkEOZUrYEsuUTLE9DUXdrNdVc
Eb0mCvWGOGvTjnkQY0uMYE2dWK5JQmcgIxEKnUaAFLv4H9GkivA0yP06SIBGCCce
qIj9Zh4S5p2EtXqi7+UxdNoXQlVCaahREJoUYHsyJLwpl0yfqQ62Kv5KZetsfegI
9pVbe+jM68kVXuHHxPPhAc1Xxh1ecmpIqG1f2HvwVta3vwzEf+PC/zD0O8xfUUah
T/Ik3MmVuQNxHKt6TUUuV6cdxui7Fz4JpglSc2BA49Usum0bBuj/dgUDGRAoqPpS
XO20JjmwFYAlCRQeMMBKQe73hmhLENoeAKRTgQ2jmtX5R7hr+073Q9ePpHZk4nFK
KXNQtfwcBH6CvO0R5NiYpbQ81GddKcqLldfNtVWqSxNU5ScvHwIkusoqZnZLbm4w
Yk1qVLxWOqfylve73cNiPzWe4OLenAX81mSGn9fWerMkrMr4B49lnPmNOeiMamdQ
HlwwD9R6nbXLcSfXRWjGt2uGhcfmApdV1q5Nt5TEYOCUrfK+RkMw4s+laIBuJOK2
qvo34dq7F0o+Xm13urjSUR7Nygun5gsHVUuynSpIHaxAgis4QqF6pAxbMSw8OyXn
OEPdnmqO2QJqN24GPpCNOeGRxG7xwBV6u1GWrURYxunR5BvmETRbfKxllXl89cjt
9aoGG8s+fv+1JVP85z/7T+ZxMxV8CQAA
`   }
}



func (resource *ResourceData) Io_beinghappy_happypulse_d_production_key() *ResourceData {
    return &ResourceData{
    name: "Io_beinghappy_happypulse_d_production_key",
    data:
`
H4sIAC9p4lYCA21Vx86raBLd8xTsUV9yuhqNRPgw0UQD9o6cwSYanr7/vrPpxdSi
pMpHVUcqMalgYV3nJt3WYoHgHynnphjz/rwnQ/EbbqZfadGMVZ283+evP/q99Uvx
V/7rvS31r/c85Vu2NtP4p7ifsqQ3ilOTf8OyAks8zOMwIcC8CNMcrMgwJsGAgzER
5mmY5f5JYASYlGBWgXkC5n48HExJMPTT41/AfsP/uU//sv8L/fWPiOCm3WHPF2DH
00IhALABnn8ikKVpYDo0URAMSXCBsNwINgr4WcXdbaf12yKR+mRg1qkNrO4+yOTL
InOcBeEjX8zRLaGILxPQ7Jwcj+GmDYR8MLGa9o5b1LgA2lIg92wBT9HM7ckYSMlY
PURuvlyJaM6SrQ70QIq+GZy2eVwLKuw3/5G/J8XKHc5Kv9u9bzljcPeAv03P9Yk2
o3NYePbTJ3TWwc8DF9IGzLE5uwVnOxZfvxEsbZC2ymGOYT+Z6uDXY2TpvgtyA8UP
wzYspK825Mkfq+cUjgcxbUnMR+aqKH0ytdtJgt3HqNtfJVLW3ssCRYs7Q2/Ob0tT
QZHIkXDiY32nP7UJVuuC3Nkl6W+l5k5JyTN4CPusXGJ6mcW4ShxYyZV4J4cmC64g
CtPPsiX3oqRY05OlAt8fPr1uwVkpoWA+8pfNRxit6DZV2ENGdZ23JHUZkcHd98Ks
Xb6zo7yA9SEM9uiXITbx+FVDZXL7Pm3kw2aW2NF17YRM1iCKNdP8ZvMTSRe5vkWv
R+J46WiE7Ns9H89VNJ0MQ1dhvaCrXF9s7qDVS4xPqWnziDqWnusx89q1ZpO1y5d2
/LPGG/e5cWATj4qj653OKL0mi+oJvSx6O0QjvyvZnNwfgyktJzpyKUH2GBEp/a0s
wdVbTdqGSBREqf+VovddKYMp9jBVECDK7m7n1VkfIbt0W6viJ7CIMxhewsutL3wU
1aW83PzAOswmTMWXRB3d4swjPQTbbISFQpbdXxiQqicQcC7AYh/JBfOQsJ/TJDeG
Y3Y6Viq6pGjEc+dqYyRdNMXKTNmOC0eCgwgUCRBZk8poPU7bTfJDdM6Oe017QtFy
sbtJJH47d/KuNCp/5iQfEttE36wfd3HpthdUSWhIUaCb4jmNN/Xs5bovojedTYsf
2GlFdMSCKyvrqYeyMu1iy1nnW9kfxFhoMpC8JUBX36GLhYA0gvFRm9+k7I4xIZ6v
ZaWlMqB3nQgzASwkr46rT9+CumK28k2gc0JAidxSskwlp9ZvC2769cO8nhoVlUPs
sFk+LlfsYgzI3pPXjP1KSVS8zk6YYCDqdC7JoE+ZNmPsadcJRJUptEIhIkxXjspX
Qqt+v4e1VdUW+wN5mRfLVCUyrvPSE5qaJ9EUSnkZv99aHj+lthbRufqKGFNIoaDv
33tbsi0IEKdNiC+2ma8BB8TqVS7/fIr4Fd0dXIUS5Y1iYEHlOlH3jbsvs5QZ/tZ2
W7Bl9O2FTEf/zMZedaaIwlPt9M/EKNyS4Vs8tw4EYtnnyi8UN81WGIT1yxR23aaX
H8higX3QzhXrJcgZkVuyWDbP3I6KAFTZTE8I6z4ACxXxU7PITzSMt283XFk6qsw+
cSmla6Smky5jErhcYS2ykg1bjeUXfGfCUccJ5BNPOQaU7TQSI5H/CeI2fx8+CaRh
8Vgj6VRj2p+9cRsdoozEcR/T0FOUeyRgyUnJmEsgqFEZ0D3sEM01xEp6VLNSNC9h
svUhmY0k0zUccOnOY8RGVewD0aTeP9Qb81QdzmfbvSHteoUSPW2NHqDS6i2x7LKf
B0mTKysxQoTeDRA/zE998LVZvevadHysffM/7PghOkHjmG8hEB2jM91t5CLUYyYc
BG5i0z5tZ/Kgso46pCbcmZew+LorEfd9ONjE9//3bMBd/v+v5m8IBAhNRAcAAA==
`   }
}



func (resource *ResourceData) Io_beinghappy_happypulse_d_production_pem() *ResourceData {
    return &ResourceData{
    name: "Io_beinghappy_happypulse_d_production_pem",
    data:
`
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
`   }
}



func (resource *ResourceData) Io_beinghappy_happypulse_development_key() *ResourceData {
    return &ResourceData{
    name: "Io_beinghappy_happypulse_development_key",
    data:
`
H4sIAC9p4lYCA21Vx66sCg7c8xXs0Qy5gat5I5GhyTSp2UEDTc7569+ZezdvMbWz
bNmlclnmki/IrutcpduaLwD4g2Ku8j5rLzPp8l+gkozjZW/tkv/b3pbyd0U7fJJW
yy9V+AVKAsgwIMWAPA+iAoiJII+CFAGSLMg/QJIBEQFEJZAnQIQFBQkUERCnQYkB
cQoEfnr8Y/ov8D/m8I/4v8C//gdOlFUTdF8saLtqwHoiqInv3xnAUFVxZFWOZTWe
dUT2jKVVtVMo80aOaFHYNkrpYg58bbcev4mYhxqdF0u85GcrzBcbSGB0yRzxliqW
va8aYZBhZy9ZoxHHqQlPh/1kVjJLbCg4TdOTr5hPRk73vXaJms36AdiBjCvR+Nr5
DS4Tm/16uJmm+f3OO70oQ1bG5I5WTiYwbXO4N6hXPicyEsU2TihzPimACl+y4uza
aHwNjtFfSx9A3nofcDZBCFoVrcFf77PTNa4/7q7Lu3MwDdeX3wt1l+mzAZI+DCuz
ShgEE796xlvrxJgdOhWfGpu9AoZh0Q8ns+kYY28kU9OHcy3MAxk3+3WpSgEUJm0K
U/p5m0tysPomFQK5lE3v8zGWwAPbbwNZHarAOizHDj9iSyK5WY4/sY2mDUcHxHL7
poKmog1PTYojfz3wmMgZCWOQx85t1CWtWRt1kqtvHmWIEBoVwQmTXM1ReG9mDfCN
2gJVhm3nDL4uXHv5nsZTp05cFMp4bruAoTIomRucpJjozfniKDrSIxBVb//yemcC
fYA1lYCHtM7uMbcTJh8WGLlT5AA5yJJoqtmkRtqj9IQKVDx2/gUjFrUztixdB6oO
gFb7Q3xvCvXR93TMB0QOqrdQn5u74FCfcqWKOXdsqjUBh7Ur8al/6Gp+V9Nn7yEl
44ASzVtigVCXyagMRSrRN1eFMIxPymPfdMXPzKFSd8Zm6/DCqy62GaF8mjtWiMam
KjUBDettnxX571tkyZLzfohTw2vBoykTYyycyuc6DnPXSq8cG3lxcGvBjAw94Txf
jQalB5DW3R9t1zZnBg/yjnwPqsLcnC/Vd+E/6WZXHxQezeErdJ7QGQtJRtDa3ihj
6U7m/cwBRq7t4imJYV9xeYmReXpdzvsR2WEkz9aMurFsVDgz3L3+VYKxKbfQwh5/
GGNMJgJl7OlH1I3prN7b+g1/HD+LbAqdSfZY02BpfQjlewxO0fpa+rByLFfVZzJI
icibM+LnmBqWO0bNZQynoL2apnfEFW7EVc9PK7sJGvGs3fh0rZwZpDyJ2ZqapVj4
j3mTc6dFDwDbKb2InOlhRjT/bqJJwG9txlu/n7cKedo/+8HjZPlNuXp9uku9n1z6
TtHGbts0eQMupkRO0Xb9cQgPvvFN3PG7lIx+nEzH6uK622ueiOOS0FdDY+yQjpLu
kVcVPxf15FsKwK+ALWqjwnTlDYU+FiHCPKQz+x2dTKrT2jEjTxDUkGekKD2tIH7t
ivrqoYHGNiheOYAPSRfz4LYPCL5LkX23bbc2/qh8SUa29zllM3QjPXubJPOwgdNs
1wZPU5ZoRPUMQCq/2p9QUZhcvGO6/ajh5JiDHV30lmOZLX8uxk7kGTOZCwLvJ48+
1gA6sV4YdsaCZiCNRrP1X9Dorx9W65tauU7qqyntaelavTDKlEsW11q6aoxWt5xq
T41s+Nho+LYtruYB6ToZiv/8UGadvZa+sBpuQQTLUtxnbnJISi2+b7ZMKEa3fNL/
pHiXwjVbsR5BQEpBAOvevuYXzKGLg7DGho/b8ukR+aSwaVTcr7C5eegUzhtTpQcK
ZXwe1xbCNW0VPOTqgTXARyx1tx9aWqTxQZNJif9mtPOBeXqX3qG2SaXdmgYeCJHq
eppH1Pxuq9+//vrzcERT+P/v5m8/985iLQcAAA==
`   }
}



func (resource *ResourceData) Io_beinghappy_happypulse_development_pem() *ResourceData {
    return &ResourceData{
    name: "Io_beinghappy_happypulse_development_pem",
    data:
`
H4sIAC9p4lYCA6VVSdOqSBbd8yvYG/1EcIAX8S2SOUFQJhl2TAKaiAKS4q8v8Ov3
uiqquzZNBEPevHmHzHMOfFyQoO/bKnn2eUeQ03Vuq/yWodGM6/wnCe53lJNiPuSo
udf5rSfhwSGPz64knbwdqjTvfpJV8yPJq1tRxvf7+OPzvD9Rl3/ioSaNkZ6PUPxJ
yiLJceSOIwWBXIkkLZHCitytyQ0ghS254UhKJFcyKaxJCpCiTEoUybCkzJHMjiS6
Z3LJ0/5r6UHx63+lXArm1/9V8/LgfbEazcoKbQvcUvjyHKLqumfefn0Gy8O/48Nb
+mN2/h75TYsyXGW/8+Ytaeco7qvm1v2nqH90I4W87atzlX6GJHj2ZdNW/Uj8a754
SYEmKUi2C2UoAFf6WAkDQhm5gsCrdAEw5EEBoYPCBJWcoHASFq1Q05sIlkNqAkuS
eQvgIgpfwhtofGGeCB6ELkAn17ANLFmheLIsXQR3PlX5LnJ4LamNp+FgrBefub0o
mL/mgoTROiJSs3ukeIWtnOjIfw2pchqhrKFEkanY557pW5YMCSpg5UkA450V8DiZ
/E80NyaKzRCxbyMo2SirT13C8CiteGcKVGYKGpKbUVj0aZy+69g3v20Vz2eB3UzZ
75l6xYRapqbhevjgSivDvY6maLz9jy2cbdRv20XgkSg1hhAqQlE7WlO7bPW6EtPG
GMoUeV9rKPY3t1iRcapen99vhDP/9Y5cWTEk61cb3K825raJ775XKLlZheNx0xHx
q5QuCm8uvUaXKDC28B8SEN8ZJNeYjnnOIEwbfpCa8SBFqoE0cWpjb4DrZ44vDeF0
MnBRSJUBKEVwHgrhwIQRLYkHlgfAGvIiBrODDpoJE5YoX32q4IbN6nBVqyBg2VEK
9oyY0/72fD6W1z2xtTroqLmkd5toi7l26b9WwDV8CMSjfnQZp2JFGLZvmxeOpTFw
9qPLqLXjDlzfD5K+J1Yvpm6PBzpqSrobAOclaZYjndV3C7lhBf6Yrvp+Z5bbsF/I
IUjCOMFndeVRLH057LYpkb1V51YutnUbMNJt6CMYZgL7uI4Aq1Svj1mwOTBHkaPb
netrCw/vtYsSdo1yFg9xtE0IpaBeshVV7+Egd7X9yKIY0lrLuKEDegpu6ZaiE//O
jce4N5fLJWvbcUM7CTUuLI9qnsRFWbGM6NRrbg+LjGMZTB3u/augXy0Az/WJPdaj
4mSWdovaRSHk20u9B4XBA6BcikJFMxv5tQuymVmqtZbkwvKWEdq8eeqJAzt+oaup
spbOLMB64aZYmAlnUy4PIAYiOBOfhY4hKSLwC96B2ivVHkkYhnDotruWtlc32UPO
pX8LArDYGQ8ZBM70Hc7HLUwVqEu+uD7Ka6VwtHnlLQkX5+lODUUQOmXCh8xjKBhQ
oUWovL2JdLEvP0PaKxJ6XRCZUt7TkZ8gq030je4hLVMT16vc4cvkdi1SRR4n1hVh
YL4zf4Um3zKkTZSqdpnUJiKgwtVQtZtpERUrpyecMa685kDVTJiUmR3tMq2tIpu0
IgmMIvQ3V6iYw2S/E794n9BhkQUm2gt/qwYn9BSQuf5aiNKbfY8mooXBFGDuI1W1
MmTs+9zXnHHK1Ef+hkrHNdatEOo45HnLU0Eh+apfUpkKtvuRYzImfRLhN7enTeH6
Pf0pH4W0NBgS9dlxFduzFmDrwYNG3CpH5XeAaxREiEgUDkcBfP410F/LTOhN+ScZ
GzImu6a1We5rcyQSEey/QcROnDYEfg2whD9gQfwkP0LxpxYwxBII9W+Zzy4RCMVJ
0ixh1vr/ov/iBBQeKIga1Pr1oIuoWbzUwisnXLpXdDuVg99jYhLEs663LEwc/XFc
q5MGqK/lwq8u9TisY8nPa0CnzzNDw/37jCS1ufa4yHByealn2g6IXhslp4Z9+xB7
3G2k+G0eTs5JbNKnUu7UY5DbE57esZ6V3L7d6W0YP9uXsbmnG0PNeY+oYOsnbS8i
t1ulx+VroExmSb08ZhkXZzreRuXmrba1ZjVe1sb2Prg7i3ZsEEVZdFDrb+J0Ut94
4I59dtJt+nguDPY8RHtxYecrl76dVw+o76uHRd3oU7bUy2W1eYRsvdSV/QVogUoc
bi0vJShemCfRuu52sgXTUFzlgDlekKZJeWlRfg6NDm+e0nrpjpvVma/Zr+8/tmSK
f/9f/wHuwRHAcgkAAA==
`   }
}



func (resource *ResourceData) Io_beinghappy_happypulse_production_key() *ResourceData {
    return &ResourceData{
    name: "Io_beinghappy_happypulse_production_key",
    data:
`
H4sIAC9p4lYCA21Vx860iBG88xTckZc8A7+8lohDzgzhRoYhpyE8/X7+97KWXLdW
l1qlVncVm1Qgs21Lk+5bsQLgD8qlKYa8u4ykL36BUjJNl7V3a/GHtYz5H9a+1r9p
3ZglnVpcMv8LxFCQwUAUBTkUJHBQZEACBWkBZGmQJ0AKB58PkGJA8gmKLIj8Jj8w
kBRBAQOBnxn/kPAL/Lcx/qP+D/Cv/4IVXrIBOi4DWo78ZjwBVIXodwfQZVmYGJll
GJVjbIG5MNfkD/nNtncOzYvjuFmelJOevGWyTna3eXP3p9adl1db0K6NABFXpb9k
r6ycvTjURxOBPoXZ03joPmkn/mypHYzmjgp6E13amVCTZ8iqa8Ubqx4+1gLfj8Yt
pAA5R0qpOqq3S0+IOaI9Ttm5YStfjmyXfb9u18yCEukm6ev1VVwNZzzuoRklEFOP
VqWdnuW9Y2CxiVcWznICSNLysi12HuEk5EoQN0VZy4mwlPvGhTI6Phez1fzoGoBY
jiYkl8UoEb8+ipJLiFijZoVEumatD5JuOoxYRRxR6CJeVtnKD5YsuyaAIejythio
YCb9tKhRUaclpIhpPm10o29pIelkmlcoUdyyPmSesRmWGX+WzbrBdKvT7q8pZDgj
ADWjhcds8YXEOsF8rNlvq5AwyVKkBMKwih/TsPHt1Vh5k5w5KghtvInm8x7Zwl/Y
GaizLiPwgJsmGvLWz/1i+4favcRuT6TkGlRJ8PWmuyzPbyx/0HvS2jbHp2r7DIKP
uqqAOTqwSspQd3d75ATxw+c2kyET5AwYj2iJ5qgE5TEnAkOt1lkktRJUy7arHxeW
ddhOgY+FblPoiqzgVxSWO52bIfca9pg0DoyclGE62cz76BzEcEjtfT2pjaoV/kIJ
hG7vsAf2Ik6Goabj4saCz6zSlXiRAwdDbxG95ozrr6pskE2Mzk4KaMMnrzxVAo0r
xas1jjQD/H5XVFLgqkhgaLyoDThenqPXB98PYdr58pn2AeKu11MQqA01AntdmiX9
LtXi36hH50ANv2m7ftaNHlSP4SJhDm3D4EwicUIIe2rrp35oltS0nfqRb+c6nfCp
btBC1l5DcS8E8Az8sqZWEDsruBUu6xCEEMSsXd28lXFS38ZPnDzXE6sG7u0LMyFt
1yfwfytG+rIBDLw7mw+0ccqsdR6hqowXD8LoiyyDy1+5CHDnu9Wba4W6laTzEp2x
i0pvVzZE4cIyHbCVmr18+wsZDKGONKNFF6tGMzHeH/6zo8tr4YoCo6yDeuZtm1JL
mEOTYY9NP/SEUmLAlye2nuZL7mnbniwUHVnEERwp3LZkd5CkCa/olLT+SGay0T3x
+Xt1U0wSGMQvF4s/geQxb1LU9yWBtQv8EnCNkclw4Zeeem2rRYfaFgpWoe6+8hg6
eoPQadDk8lCzdMCa4QZsZ9D2d/UUQna2BKEhYSIjkh9T0daqMYNT3J33lshncgbj
bE3+vR73RIkyHZmXRzg+UOXuu2WvIjgee9O9lPb7CHdbZSubQxwmgZsRnmI85QMf
LTeEc4MNncWt36Gtd4Om4gDVyuxHFvCvTlGsUUnUxE976koalFdrKJrTqB5KXTjZ
8Rir0UOON9ryp3bjOirFbWUCkhmlvqlr2E5Xp7mgIX21euztbkw36s1m59trX0Rm
NBv1hdXDVL/5p2jTm9lVXXnhJeCtAlYdf0u2oTFyR+1+7T/3hNHty+snMfmK8fCc
1SH3wrNSuinnU2NIcQ9mEBQygdIgpzx8ZhmErrPltPotWfNhW2s9B14U0MgXqcPF
3ZhjdvwS89SreVRwL9H1FFf4gAEzKg74tv78j3fRuAx5ESL+GN/ZmPS+w0lxkPHC
DwS6NcehwDoDMcOKMn/++XfgCAb//+OG/d/E/Av9neUdQQcAAA==
`   }
}



func (resource *ResourceData) Io_beinghappy_happypulse_production_pem() *ResourceData {
    return &ResourceData{
    name: "Io_beinghappy_happypulse_production_pem",
    data:
`
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
fY/M7+/Piy0awj/f6z8AUfEAfXAJAAA=
`   }
}



func (resource *ResourceData) Server_ca_cert_pem() *ResourceData {
    return &ResourceData{
    name: "Server_ca_cert_pem",
    data:
`
H4sIAC9p4lYCA2VUubajOBDN+YrJffoYDNgQTCBAgNhs8dgzNoPZjBfWrx/cr5N+
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

