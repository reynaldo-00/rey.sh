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
            // const resp = await axios('http://https://rey-portfolio.herokuapp.com/repos');
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
        background: rgb(29, 29, 29);
        color: white;
        padding: 48px;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: flex-start;
        padding-top: 100px;
        padding-bottom: 100px;
        h2 {
            letter-spacing: 14px;
            font-size: 30px;
            margin: 0;
            margin-bottom: 20px;
        }
        .carousel {
            width: 100%;
        }
        @media (max-width: 620px) {
            width: 100vw;
            .carousel {
                width: 100vw;
                padding: 40px;
            }
        }
    }
</style>
