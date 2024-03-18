const url = window.location.href;
const articleId = url.split("/")[4];

fetch("/api/fetchOneArticle/" + articleId)
.then(response => response.json())
.then(data => {
    var div = document.createElement("div");
    div.classList.add("border", "border-1", "rounded", "rounded-3", "p-4", "shadow", "bg-body");
    div.innerHTML = 
    "<h1 class='h1 text-braek'>" + data["article"].Title + "</h1>" + 
    "<h6 class='h6 text-muted'>Created at: " + formatTimeString(data["article"].Created_at) + "</h6>" + 
    "<h6 class='h6 text-muted'>Updated at: " + formatTimeString(data["article"].Updated_at) + "</h6>" + 
    "<hr>" + 
    "<h5 class='h5 text-braek'>" + data["article"].Content + "</h5>";

    document.getElementById("data-container").appendChild(div);
})
.catch(err => {
    console.error("Error: " + err);
});

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

function handleDelete(){
    const currentURL = window.location.href;
    const deletedArticleId = currentURL.split("/")[4];
    
    fetch("/api/delete_article/" + deletedArticleId, {method: "delete"})
    .then(response => response.json())
    .then(data => {
        var redirect = data['redirect'];
        window.location.replace(window.location.protocol + "//" + redirect);
    })
    .catch(err => {
        console.error("Error: " + err);
    });
}