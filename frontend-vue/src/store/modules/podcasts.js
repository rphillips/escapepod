import Axios from 'axios';

const state = {
  podcasts: [],
  episodes: [],
  podcast: {},
};
const actions = {
  listPodcasts: async ({ commit }) => {
    try {
      const { data } = await Axios.get(`/api/v1/podcasts`);

      if (data) {
        commit('setPodcasts', data);
      }
    } catch (error) {
      throw error;
    }
  },
  getPodcast: async ({ commit }, podcastID) => {
    try {
      const { data } = await Axios.get(`/api/v1/podcasts/${podcastID}`);

      if (data) {
        commit('setPodcast', data);
      }
    } catch (error) {
      throw error;
    }
  },
  listEpisodes: async ({ commit }, podcastID) => {
    try {
      const { data } = await Axios.get(`/api/v1/podcasts/${podcastID}/episodes`);

      if (data) {
        commit('setEpisodes', data);
      }
    } catch (error) {
      throw error;
    }
  },
  createPodcast: async ({ commit }, feedURL) => {
    try {
      const { data } = await Axios.post('/api/v1/podcasts', {
        feed_url: feedURL,
      });
      if (data) {
        commit('createPodcast', data);
      }
    } catch (err) {
      throw err;
    }
  },
  importPodcasts: async ({ commit }, opmlData) => {
    try {
      await Axios.post('/api/v1/podcasts/import', opmlData, {
        headers: { 'Content-Type': 'multipart/form-data' },
      });
    } catch (err) {
      throw err;
    }
  },
};
const mutations = {
  createPodcasts: (state, payload) => {
    state.podcasts.push(payload);
  },
  setPodcasts: (state, payload) => {
    state.podcasts = payload;
  },
  setPodcast: (state, payload) => {
    state.podcast = payload;
  },
  setEpisodes: (state, payload) => {
    state.episodes = payload;
  },
};
const getters = {
  podcasts: (state) => {
    return state.podcasts;
  },
  podcast: (state) => {
    return state.podcast;
  },
  episodes: (state) => {
    return state.episodes;
  },
};

export default {
  state,
  getters,
  actions,
  mutations,
};
