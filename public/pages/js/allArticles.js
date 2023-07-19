const delete_btns = document.querySelectorAll('.btn');

console.log(delete_btns);
delete_btns.forEach( delete_btn=> {
    delete_btn.addEventListener('click', () => {
        const articleId = delete_btn.dataset.articleId;
        console.log(articleId);
        deleteArticle(articleId);
    });
});

async function deleteArticle(articleId) {
    const response = await fetch(`/articles/delete-article/${articleId}`,{
        method : "DELETE"
    })

    const data = await response.json();
    console.log(data);
    if(data.success){
        result = confirm("Are you sure? If you delete this article, you cannot recover it");
        if(result) {
            console.log(result);
            document.getElementById(`container${articleId}`).remove();
        }
    } else {
        alert(data.message);
    }
}