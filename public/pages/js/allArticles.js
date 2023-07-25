const delete_btns = document.querySelectorAll('.btn');

delete_btns.forEach( delete_btn=> {
    delete_btn.addEventListener('click', () => {
        const articleId = delete_btn.dataset.articleId;
        console.log(articleId);
        deleteArticle(articleId);
    });
});

async function deleteArticle(articleId) {
    result = confirm("Are you sure? If you delete this article, you cannot recover it");
    if(!result){
        return;
    }
    const response = await fetch(`/articles/delete-article/${articleId}`,{
        method : "DELETE"
    })

    const data = await response.json();
    console.log(data);
    if(data.success){
        console.log(result);
        document.getElementById(`container${articleId}`).remove();
        document.getElementById(`delete_btn${articleId}`).remove();
    } else {
        alert(data.message);
    }
}
