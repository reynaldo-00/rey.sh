<template>
    <div class="item-container">
        <h3 class="title">{{title}}</h3>
        <img class="preview-img" :src="cover"/>
        <p class="description">{{description}}</p>
        <div class="links-container">
            <a :href="deployLink" target="__blank">SITE</a>
            <a :href="repoLink" target="__blank">REPO</a>
        </div>
        <div class="tech-used">
            <span
                v-for="(item, index) in techStack"
                :key="index"
                class="tech-item"
            >
                <StyledSVG :logo="item"/>
                <span>{{item}}</span>
            </span>
        </div>
    </div>
</template>

<script>
import StyledSVG from '../../shared/SVG.js';
export default {
    name: "Project",
    components: {StyledSVG},
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
        .tech-used {
            display: flex;
            justify-content: space-evenly;
            align-items: center;
            width: 100%;
            background-color: rgb(29, 29, 29);
            border-radius: 12px;
            padding: 20px;
            margin-top: 20px;
            .tech-item {
                display: flex;
                flex-direction: column;
                justify-content: center;
                align-items: center;
                margin: 0 20px;
                font-size: 16px;
                font-weight: 700;
                span {
                    color: white;
                    text-transform: capitalize;
                }
            }
        }
    }
</style>
