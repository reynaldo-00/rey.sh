<template>
    <section class="tech-stack section">
        <Stack/>
        <div class="info">
            <p class="description">
                Currently trying to expand my knowledge beyond my 
                proficient stack (React and Node.js)
                and create applications using Vue on the client with a server 
                side written in Golang. All while making calls to and 
                creating GraphQL APIs.
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
        justify-content: flex-start;
        align-items: center;
        padding: 48px;
        width: 100vw;
        background: rgb(29, 29, 29);
        .info {
            width: 80%;
            margin-top: 40px;
            h2 {
                text-align: center;
            }
            .description {
                text-align: center;
                font-size: 18px;
                line-height: 28px;
            }
        }
    }
</style>
