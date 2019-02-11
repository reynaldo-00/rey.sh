<template>
    <section class="section tech-stack">
        <div class="info">
            <h2>TECH STACK</h2>
            <p class="description">
                Currently trying to expand my knowledge beyond my 
                proficient stack (React and Node.js)
                and create applications using Vue on the client with a server 
                side written in Golang. All while making calls to and 
                creating GraphQL APIs.
            </p>
        </div>
        <Stack/>
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
    .tech-stack {
        display: flex;
        flex-direction: row;
        justify-content: space-around;
        align-items: center;
        padding: 48px 24px;
        height: 500px;
        width: 100%;
        .info {
            width: 50%;
            .description {
                text-align: left;
                font-size: 18px;
                line-height: 28px;
            }
        }
    }
</style>
