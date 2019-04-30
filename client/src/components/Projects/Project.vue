<template>
    <div class="item-container">
        <div class="info-container">
            <h3 class="title">{{title}}</h3>
            <p class="description">{{description}}</p>
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
            <div class="links-container">
                <a :href="repoLink" target="__blank">REPO</a>
                <a v-if="deployLink.length !== 0" :href="deployLink" target="__blank">SITE</a>
            </div>
        </div>
        <div class="img-container">
            <img class="preview-img" :src="cover"/>
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
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        margin: 80px auto;
        width: 100%;
        max-width: 920px;
        background: #0e0e0e;
        height: 380px;
        border-radius: 2px;
        box-shadow: 0 2px 3px rgba(0,0,0,.5);
        transition: transform .4s ease-in-out;
        transform: scale(1);
        @media (max-width: 620px) {
            width: 100%;
            height: 700px;
            flex-direction: column;
        }
        @media (min-width: 620px) {
            &:hover {
                transform: scale(1.03);
                .img-container > .preview-img {
                    height: 103%;
                }
                .info-container > .links-container > a {
                    opacity: 1;
                }
            }
        }
        .img-container {
            width: 50%;
            height: 100%;
            display: flex;
            justify-content: center;
            align-items: flex-start;
            overflow: hidden;
            @media (max-width: 620px) {
                width: 100%;
                height: 50%;
            }
            .preview-img {
                height: auto;
                width: 100%;
                height: 100%;
                width: auto;
                transition: all .4s ease-in-out;
            }
        }

        .info-container{
            width: 50%;
            height: 100%;
            padding: 0px 20px;
            display: flex;
            flex-direction: column;
            justify-content: flex-start;
            position: relative;
            @media (max-width: 620px) {
                width: 100%;
                height: 50%;
            }
            .title {
                font-size: 22px;
                margin: 0;
                margin-top: 10px;
            }
            .description {
                font-size: 18px;
            }
            .links-container {
                width: 100%;
                display: flex;
                justify-content: space-between;
                align-items: center;
                position: absolute;
                bottom: 20px;
                left: 0;
                padding: 0px  20px;
                a {
                    color: white;
                    opacity: .5;
                    font-size: 24px;
                    &:hover {
                        opacity: 1;
                    }
                }
            }
        }
        .tech-used {
            display: flex;
            justify-content: space-evenly;
            align-items: center;
            flex-wrap: wrap;
            width: 100%;
            .tech-item {
                display: flex;
                flex-direction: column;
                justify-content: center;
                align-items: center;
                font-size: 12px;
                font-weight: 700;
                span {
                    color: white;
                    text-transform: capitalize;
                }
            }
        }
    }
</style>
