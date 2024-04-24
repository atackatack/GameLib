let timerRandomGame = 0;

$(document).on('click', '#buttonRandom', function (event) {
    $.ajax({
        url: '/api/v1/game/random',
        method: 'GET',
        dataType: 'json',
        success: function (body) {
            clearTimeout(timerRandomGame);
            const element = $("#labelRandomGame");
            const newText = body.data.name;
            if (newText.length == 0) {
                return;
            }
            element.text(newText);
            $("#randomGame").show();
            timerRandomGame = setTimeout(function () {
                $('#randomGame').hide();
            }, 8000);
        }
    });
});
