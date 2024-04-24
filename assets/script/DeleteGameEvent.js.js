$(document).on('click', '#buttonDeleteElem', function (event) {
    $.ajax({
        url: '/api/v1/game/' + $(this).val(),
        method: 'DELETE',
        dataType: 'json',
        success: function () {
            refreshTable();
        }
    });
});