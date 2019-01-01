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
    p {
        width: 605px;
        font-weight: 600;
        font-size: 22px;
        text-align: right;
        float: right;
        margin-top: -30px;
        line-height: 28px;
    }

</style>
