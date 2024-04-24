$(document).ready(function () {
    $("#search").keyup(function () {
        value = $("#search").val();
        delay(function () {
            $.each($("#gameTable div.card"), function () {
                if (this.getAttribute("data-name").toLowerCase().indexOf(value.toLowerCase()) == -1) {
                    $(this).hide();
                } else {
                    $(this).show();
                }
            });
        }, 50)
    });
});