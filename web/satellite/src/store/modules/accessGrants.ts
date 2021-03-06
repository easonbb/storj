// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

import { StoreModule } from '@/store';
import {
    AccessGrant,
    AccessGrantCursor,
    AccessGrantsApi,
    AccessGrantsOrderBy,
    AccessGrantsPage,
} from '@/types/accessGrants';
import { SortDirection } from '@/types/common';

export const ACCESS_GRANTS_ACTIONS = {
    FETCH: 'fetchAccessGrants',
    CREATE: 'createAccessGrant',
    DELETE: 'deleteAccessGrants',
    CLEAR: 'clearAccessGrants',
    SET_SEARCH_QUERY: 'setAccessGrantsSearchQuery',
    SET_SORT_BY: 'setAccessGrantsSortingBy',
    SET_SORT_DIRECTION: 'setAccessGrantsSortingDirection',
    TOGGLE_SELECTION: 'toggleAccessGrantSelection',
    CLEAR_SELECTION: 'clearAccessGrantSelection',
};

export const ACCESS_GRANTS_MUTATIONS = {
    SET_PAGE: 'setAccessGrants',
    TOGGLE_SELECTION: 'toggleAccessGrantsSelection',
    CLEAR_SELECTION: 'clearAccessGrantsSelection',
    CLEAR: 'clearAccessGrants',
    CHANGE_SORT_ORDER: 'changeAccessGrantsSortOrder',
    CHANGE_SORT_ORDER_DIRECTION: 'changeAccessGrantsSortOrderDirection',
    SET_SEARCH_QUERY: 'setAccessGrantsSearchQuery',
    SET_PAGE_NUMBER: 'setAccessGrantsPage',
};

const {
    SET_PAGE,
    TOGGLE_SELECTION,
    CLEAR_SELECTION,
    CLEAR,
    CHANGE_SORT_ORDER,
    CHANGE_SORT_ORDER_DIRECTION,
    SET_SEARCH_QUERY,
    SET_PAGE_NUMBER,
} = ACCESS_GRANTS_MUTATIONS;

export class AccessGrantsState {
    public cursor: AccessGrantCursor = new AccessGrantCursor();
    public page: AccessGrantsPage = new AccessGrantsPage();
    public selectedAccessGrantsIds: string[] = [];
}

/**
 * creates apiKeys module with all dependencies
 *
 * @param api - apiKeys api
 */
export function makeAccessGrantsModule(api: AccessGrantsApi): StoreModule<AccessGrantsState> {
    return {
        state: new AccessGrantsState(),
        mutations: {
            [SET_PAGE](state: AccessGrantsState, page: AccessGrantsPage) {
                state.page = page;
                state.page.accessGrants = state.page.accessGrants.map(accessGrant => {
                    if (state.selectedAccessGrantsIds.includes(accessGrant.id)) {
                        accessGrant.isSelected = true;
                    }

                    return accessGrant;
                });
            },
            [SET_PAGE_NUMBER](state: AccessGrantsState, pageNumber: number) {
                state.cursor.page = pageNumber;
            },
            [SET_SEARCH_QUERY](state: AccessGrantsState, search: string) {
                state.cursor.search = search;
            },
            [CHANGE_SORT_ORDER](state: AccessGrantsState, order: AccessGrantsOrderBy) {
                state.cursor.order = order;
            },
            [CHANGE_SORT_ORDER_DIRECTION](state: AccessGrantsState, direction: SortDirection) {
                state.cursor.orderDirection = direction;
            },
            [TOGGLE_SELECTION](state: AccessGrantsState, accessGrant: AccessGrant) {
                if (!state.selectedAccessGrantsIds.includes(accessGrant.id)) {
                    accessGrant.isSelected = true;
                    state.selectedAccessGrantsIds.push(accessGrant.id);

                    return;
                }

                accessGrant.isSelected = false;
                state.selectedAccessGrantsIds = state.selectedAccessGrantsIds.filter(accessGrantId => {
                    return accessGrant.id !== accessGrantId;
                });
            },
            [CLEAR_SELECTION](state: AccessGrantsState) {
                state.selectedAccessGrantsIds = [];
                state.page.accessGrants = state.page.accessGrants.map((accessGrant: AccessGrant) => {
                    accessGrant.isSelected = false;

                    return accessGrant;
                });
            },
            [CLEAR](state: AccessGrantsState) {
                state.cursor = new AccessGrantCursor();
                state.page = new AccessGrantsPage();
                state.selectedAccessGrantsIds = [];
            },
        },
        actions: {
            fetchAccessGrants: async function ({commit, rootGetters, state}, pageNumber: number): Promise<AccessGrantsPage> {
                const projectId = rootGetters.selectedProject.id;
                commit(SET_PAGE_NUMBER, pageNumber);

                const accessGrantsPage: AccessGrantsPage = await api.get(projectId, state.cursor);
                commit(SET_PAGE, accessGrantsPage);

                return accessGrantsPage;
            },
            createAccessGrant: async function ({commit, rootGetters}: any, name: string): Promise<AccessGrant> {
                const accessGrant = await api.create(rootGetters.selectedProject.id, name);

                return accessGrant;
            },
            deleteAccessGrants: async function({state, commit}: any): Promise<void> {
                await api.delete(state.selectedAccessGrantsIds);

                commit(CLEAR_SELECTION);
            },
            setAccessGrantsSearchQuery: function ({commit}, search: string) {
                commit(SET_SEARCH_QUERY, search);
            },
            setAccessGrantsSortingBy: function ({commit}, order: AccessGrantsOrderBy) {
                commit(CHANGE_SORT_ORDER, order);
            },
            setAccessGrantsSortingDirection: function ({commit}, direction: SortDirection) {
                commit(CHANGE_SORT_ORDER_DIRECTION, direction);
            },
            toggleAccessGrantSelection: function ({commit}, accessGrant: AccessGrant): void {
                commit(TOGGLE_SELECTION, accessGrant);
            },
            clearAccessGrantsSelection: function ({commit}): void {
                commit(CLEAR_SELECTION);
            },
            clearAccessGrants: function ({commit}): void {
                commit(CLEAR);
                commit(CLEAR_SELECTION);
            },
        },
        getters: {
            selectedAccessGrants: (state: AccessGrantsState) => state.page.accessGrants.filter((grant: AccessGrant) => grant.isSelected),
        },
    };
}
