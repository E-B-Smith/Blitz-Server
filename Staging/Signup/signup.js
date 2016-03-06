function FormatPhoneNumber(phoneIn) {
    'use strict';
    if (phoneIn === null) {
        return "";
    }

    //  Strip non-numbers --

    var phonearray = phoneIn.match(/[0-9]*/g);
    if (phonearray === null) {
        return ""
    }
    var phone = "";
    for (var index = 0; index < phonearray.length; index++) {
        phone += phonearray[index];
    }

    //  Format number --

    if (phone.substr(0,1) == 1) {
        phone = phone.substr(1)
    }

    var result = "";
    if (phone.length == 0) {
        result = "";
    } else if (phone.length < 3) {
        result = "(" + phone;
    } else if (phone.length < 6) {
        result = "(" + phone.substr(0, 3) + ") "
        result += phone.substr(3)
    } else {
        var rest = phone.length - 6;
        rest = Math.min(rest, 4);
        result = "(" + phone.substr(0, 3) + ") ";
        result += phone.substr(3, 3) + "-";
        result += phone.substr(6, rest);
    }
    return result;
}
var lastnumber = ""
function validatePhoneNumber(e) {
    'use strict';
    var phonefield = document.getElementById('phone');
    var val = phonefield.value;
    if (val !== null && val == lastnumber.substr(0, lastnumber.length-1)) {
        lastnumber = val;
    } else {
        lastnumber = FormatPhoneNumber(val);
        phonefield.value = lastnumber;
    }
}
function browserPlatform()
    {
    'use strict';
    if (navigator.userAgent.indexOf('iPhone') !== -1 ||
        navigator.userAgent.indexOf('iPad')   !== -1 ||
        navigator.userAgent.indexOf('iPod')   !== -1)
        { return 'iOS'; }

    if (navigator.platform.indexOf('Android') !== -1 ||
        navigator.platform.indexOf('Linux')   !== -1)
        { return 'Android'; }

    if (navigator.platform.indexOf('Mac') !== -1)
        { return 'Mac'; }

    if (navigator.platform.indexOf('Win') !== -1)
        { return 'Mac'; }

    if (navigator.platform.indexOf('Linux') !== -1)
        { return 'Linux'; }

    return 'Unknown';
    }
function updateMessageBox()
    {
    'use strict';
    setTimeout(function() { window.scrollTo(0, 10); }, 100);
    var message =
        '<b>iOS Required</b><br><br>Open this page on your iOS device<br>to enroll for BeingHappy.';
    if (browserPlatform() === 'iOS') {
        message = '<b>BeingHappy Enrollment</b><br><br>' +
            'Enroll your iOS device for the BeingHappy trials.<br><br>' +
            'After completing this form and installing our app profile, ' +
            'you will be sent a link to download the app.'
    } else {
        var formdiv = document.getElementById('form-div');
        formdiv.style.display = 'none';
    }
    var messagebox = document.getElementById('message-box');
    messagebox.innerHTML = message;
    }
window.onload = updateMessageBox;
