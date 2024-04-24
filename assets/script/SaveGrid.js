function aggregationImg(date) {
    var formData = new FormData();
    formData.append('image', date);
    return formData
}

async function saveImg(date) {
    if (date) {
        $.ajax({
            url: 'minio/image/',
            method: 'POST',
            data: aggregationImg(date),
            async: false,
            contentType: false,
            processData: false,
        });
    }
}