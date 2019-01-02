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
        <p>{{ bio }}</p>
    </div>
</template>

<script>

export default {
    name: 'NameBio',
    data: () => ({
        first: 'CHRISTOPHER',
        middle: 'REYNALDO',
        last: 'CEDENO',
        bio: 'Software Engineer based in NYC. Currently trying to expand my knowledge beyond my proficient stack (React and Node.js) and creating application with Vue on the client with a server side written in golang. All while making calls to and creating a GraphQL API.'
    })
}
</script>

<style scoped>
    .shape {
        shape-outside: polygon(0 0, 100% 0, 100% 20%, 80% 20%, 80% 60%, 60% 60%, 60% 95%, 0 95%);
        clip-path: polygon(0 0, 100% 0, 100% 35%, 80% 35%, 80% 70%, 60% 70%, 60% 100%, 0 100%);
        /* background-color: rgba(255,0,0,0.5); */
        float: left;
        width: 500px;
        height: 210px;
    }
    h2 {
        font-size: 70px;
        font-weight: 800;
        margin: 0px;
        line-height: 70px;
    }
    p {
        font-weight: 600;
        font-size: 22px;
        line-height: 28px;
        margin-top: 39px;
    }

</style>
