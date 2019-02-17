<template>
    <section class="tech-stack section">
        <h2 class="title">TOOLS</h2>
        <Stack/>
        <div class="info">
            <p class="description">
                Currently trying to expand my knowledge beyond my 
                proficient stack (<StyledSVG logo="react"/>React and<StyledSVG logo="node"/>Node.js) 
                and create applications using<StyledSVG logo="vue"/>Vue on the client with a server 
                side written in <StyledSVG logo="golang"/>Golang. All while making calls to and 
                creating <StyledSVG logo="graphql"/>GraphQL APIs.
            </p>
        </div>
    </section>
</template>

<script>
import Stack from './Stack';
import StyledSVG from '../../shared/SVG.js'

export default {
    name: 'TechStack',
    components: {
        Stack,
        StyledSVG
    }
}
</script>

<style lang="scss">
    section.tech-stack {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        padding: 48px;
        width: 100vw;
        background: #0e0e0e;
        text-align: center;
        h2.title {
            letter-spacing: 14px;
            font-size: 30px;
            margin: 0;
            margin-left: 14px;
        }
        .info {
            width: 80%;
            margin-top: 76px;
            text-align: center;
            @media (max-width: 620px) {
                width: 100vw;
                padding: 20px;
            }
            h2 {
                text-align: center;
            }
            .description {
                text-align: center;
                font-size: 20px;
                line-height: 40px;
                max-width: 920px;
                margin: 0 auto;
            }
        }
    }
</style>
