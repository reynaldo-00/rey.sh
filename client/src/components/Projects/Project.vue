<template>
    <div class="item-container">
        <h3 class="title">{{title}}</h3>
        <img class="preview-img" :src="cover"/>
        <p class="description">{{description}}</p>
        <div class="links-container">
            <a :href="deployLink" target="__blank">SITE</a>
            <a :href="repoLink" target="__blank">REPO</a>
        </div>
    </div>
</template>

<script>
export default {
    name: "Project",
    props: [
        'title', 
        'description', 
        'cover', 
        'repoLink', 
        'deployLink',
        'techStack'
    ],
}
</script>

<style lang="scss">
    .item-container {
        display: flex;
        flex-direction: column;
        justify-content: space-around;
        align-items: center;
        .title {
            font-size: 22px;
        }
        .preview-img {
            height: 320px;
            width:  auto;
        }
        .description {
            font-size: 18px;
        }
    }
</style>
