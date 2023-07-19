const deleteBtns = document.querySelectorAll('.deleteBtn');
const article = document.getElementById(`article`);

deleteBtns.forEach( deleteBtn=>{
deleteBtn.addEventListener('click', () => {
    const articleId = article.dataset.articleId;
    const commentId = deleteBtn.dataset.commentId;
    console.log("commentId",commentId, "articleId",articleId);
    deleteComment(articleId, commentId);
    });
});

async function deleteComment(articleId, commentId) {
    const response = await fetch(`/articles/${articleId}/delete-comment/${commentId}`,{
        method : "DELETE"
    })

    const data = await response.json();
    console.log(data);
    if(data.success===true){
        console.log(commentId)
        document.getElementById(`comments${commentId}`).remove();
    } else {
        alert(data.message);
    }
    console.log("mama", data.success, "papa", data.message);
}