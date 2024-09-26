let timerRandomGameFav = 0;

$(document).on('click', '#buttonRandomFav', function (event) {
    $.ajax({
        url: '/api/v1/game/random/favorite',
        method: 'GET',
        dataType: 'json',
        success: function (body) {
            clearTimeout(timerRandomGameFav);
            const element1 = $("#labelRandomFavGame");
            const newText1 = body.data.name;
            if (newText1.length == 0) {
                return;
            }
            element1.text(newText1);
            $("#randomFavGame").show();
            timerRandomGameFav = setTimeout(function () {
                $('#randomFavGame').hide();
            }, 8000);
        }
    });
});
