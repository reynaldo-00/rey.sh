<template>
    <div>
        <div class="shape">
            <h2>
                {{ first }}
            </h2>
            <h2>
                {{ middle }}
            </h2>
            <h2>
                {{ last }}
            </h2>
        </div>
        <p>
            Software Engineer based in NYC. 
            Currently trying to expand my knowledge beyond my 
            proficient stack (<StyledSVG logo="react"/>React and<StyledSVG logo="node"/>Node.js) 
            and create applications using<StyledSVG logo="vue"/>Vue on the client with a server 
            side written in <StyledSVG logo="golang"/>Golang. All while making calls to and 
            creating <StyledSVG logo="graphql"/>GraphQL APIs.
        </p>
    </div>
</template>

<script>
import StyledSVG from './SVG.js';

export default {
    name: 'NameBio',
    components: {
        StyledSVG
    },
    data: () => ({
        first: 'CHRISTOPHER',
        middle: 'REYNALDO',
        last: 'CEDENO',
    })
}
</script>

<style scoped>
    .shape {
        shape-outside: polygon(0 0, 100% 0, 100% 20%, 80% 20%, 80% 60%, 60% 60%, 60% 95%, 0 95%);
        clip-path: polygon(0 0, 100% 0, 100% 35%, 80% 35%, 80% 70%, 60% 70%, 60% 100%, 0 100%);
        float: left;
        width: 500px;
        height: 210px;
    }
    h2 {
        font-family: 'Montserrat';
        font-size: 64px;
        font-weight: 800;
        margin: 0px;
        line-height: 70px;
        user-select: none;
    }
    p {
        font-weight: 600;
        font-size: 18px;
        line-height: 28px;
        margin-top: 38px;
        user-select: none;
    }

</style>
