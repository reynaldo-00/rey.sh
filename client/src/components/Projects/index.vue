<template>
    <section class="section projects">
        <h2>PROJECTS</h2>
        <div class="carousel">
            <Project 
                v-for="project in projects"
                :title="project.title"
                :description="project.description"
                :cover="project.cover"
                :repoLink="project.repoLink"
                :deployLink="project.deployLink"
                :techStack="project.techStack"
                :key="project.id"
            />
        </div>

    </section>
</template>

<script>
import Project from './Project'
import proData from './data.js'
export default {
    name: 'Projects',
    components: { Project },
    data: () => ({
        projects: Array
    }),
    methods: {
        getProjects: async function()  {
            // const resp = await axios('http://localhost:9001/repos');
            this.projects = proData;
        },
    },
    beforeMount() {
        this.getProjects();
    }
}
</script>

<style lang="scss">
    section.projects {
        width: 100vw;
        height: auto;
        max-height: none;
        background: rgb(245,245,245);
        color: #020424;
        padding: 48px;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: flex-start;

        h2 {
            letter-spacing: 14px;
            font-size: 30px;
            margin: 0;
        }
    }
</style>
