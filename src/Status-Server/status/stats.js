var pageLoadTime = new Date().getTime();
var refreshTime = pageLoadTime + 60000;
var timetampString1 = "";
var timetampString2 = "";
var timestampWidth  = 0;
String.prototype.repeat= function(n){
    if (n <= 0) return "";
    return Array(n+1).join(this);
}
function updateRefreshCountdown() {
    'use strict';
    setTimeout(updateRefreshCountdown, 5000);
    var currentTime = new Date();
    var timeRemaining = (refreshTime - currentTime);
    var timediv =  document.getElementById('timestamp');

    if (timestampWidth == 0) {
        var s = timediv.innerHTML;
        var i = s.indexOf("-")
        timetampString1 = s.substring(0, i)
        var j = s.lastIndexOf("-");
        timetampString2 = s.substring(j+1, s.length);
        timestampWidth = j - i + 1;
    }

    var i = timeRemaining / (refreshTime - pageLoadTime);
    i = Math.round(i * timestampWidth);
    var fill1 = "#".repeat(timestampWidth - i);
    var fill2 = "-".repeat(i);

    var s = timetampString1 + fill1 + fill2 + timetampString2;
    timediv.innerHTML = s;

    if (timeRemaining < 0) {
        window.location.reload();
        return
    }
}
window.onload = updateRefreshCountdown;
