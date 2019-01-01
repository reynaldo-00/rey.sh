<template>
    <div>
        <p>{{ bio }}</p>
    </div>
</template>

<script>

export default {
    name: 'Bio',
    data: () => ({
        bio: 'Software Engineer based in NYC. Currently trying to expand my knowledge beyond my proficient stack (React and Node.js) and creating application with Vue on the client with a server side written in golang. All while making calls to and creating a GraphQL API.'
    })
}
</script>

<style scoped>
    div {
        width: 100%;
        height: 150px;
    }
    p {
        width: 425px;
        font-weight: 600;
        font-size: 22px;
        text-align: right;
        float: right;
        margin-top: -100px;
        line-height: 28px;
    }

</style>
