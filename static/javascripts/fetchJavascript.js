function deleteBlog(blogId) {
    fetch("/blog/" + blogId, {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
        },
    })
        .then((response) => {
            if (!response.ok) {
                throw new Error("Failed to delete blog");
            }
            window.location.href = "/";
        })
        .catch((error) => {
            console.error("Error deleting blog:", error);
        });
}
function updateBlog(blogId) {
    const name = document.getElementById("name").value;
    const author = document.getElementById("author").value;
    const tags = document.getElementById("tags").value;
    const content = document.getElementById("content").value;

    const data = {
        name: name,
        author: author,
        tags: tags,
        content: content,
    };

    fetch("/blog/" + blogId, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    })
        .then((response) => {
            if (!response.ok) {
                throw new Error("Failed to update blog");
            }
            return response.json();
        })
        .then((updatedBlog) => {
            // Redirect to the updated blog page
            window.location.href = "/blog/" + updatedBlog.ID;
        })
        .catch((error) => {
            console.error("Error updating blog:", error);
        });
}
