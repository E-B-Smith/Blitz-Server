function FormatAppName() {
    'use strict';
    var appname = document.getElementById('app-name');
    var innerHTML = appname.innerHTML;
    var regex = /(.[a-z]*)(.*)/g;
    var matches = regex.exec(innerHTML);
    if (matches.length > 1) {
        innerHTML = "<span>" + matches[1] + "</span><span>" + matches[2] + "</span>";
        appname.innerHTML = innerHTML;
    }
}
window.onload = FormatAppName;
