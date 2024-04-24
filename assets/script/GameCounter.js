function getAllGameCount() {
    let cnt = 0
    $.ajax({
        url: '/api/v1/game/count/all',
        method: 'GET',
        dataType: 'json',
        async: false,
        success: function (body) {
            cnt = body.count
        }
    });
    return cnt
}

function getDoneGameCount() {
    let cnt = 0
    $.ajax({
        url: '/api/v1/game/count/done',
        method: 'GET',
        dataType: 'json',
        async: false,
        success: function (body) {
            cnt = body.count
        }
    });
    return cnt
}