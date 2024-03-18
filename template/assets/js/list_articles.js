fetch('/api/fetchArticles')
.then(response => response.json())
.then(data => {
    data["articles"].forEach(item => {
        var div = document.createElement('div')
        div.classList.add("col-lg-4", "my-3")
        div.innerHTML = 
       "<div class='card'>"+
            "<div class='card-header'>ID#" + item.Id + "</div>"+
            "<div class='card-body'>"+
                "<p class='card-title h4'>" + item.Title + "</p>"+
                "<p class='card-content'>" + item.Content + "</p>"+
                "<p class='text-muted'>Updated at " + formatTimeString(item.Updated_at) + "</p>"+
            "</div>"+
            "<a class='stretched-link' href='/article/" + item.Id + "'></a>"+
       "</div>"
        document.getElementById('data-container').appendChild(div)
    })
})
.catch(error => {
    console.log('Fetch error: ', error)
})

function formatTimeString(timeString){
    var date = new Date(timeString);
    var year = date.getFullYear();
    var month = String(date.getMonth() + 1).padStart(2, "0"); // 因為月份是從0開始，所以要加一，然後padding讓他保持在二位數
    var day = String(date.getDate()).padStart(2, "0");
    var hour = String(date.getHours()).padStart(2, "0");
    var minute = String(date.getMinutes()).padStart(2, "0");
    var second = String(date.getSeconds()).padStart(2, "0");

    const formattedTimeString = `${year}-${month}-${day} ${hour}:${minute}:${second}`


    return formattedTimeString
}