$(document).on('click', '#favoriteElemButton', function (event) {
    $.ajax({
        url: '/api/v1/game/reverse/favorite/' + $(this).val(),
        method: 'PUT',
        dataType: 'json',
        success: function () {
            refreshTable();
        }
    });
});
