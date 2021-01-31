enum Permission {
    FULL = 128,
    UNUSED2 = 64,
    UNUSED1 = 32,
    DELETE_CHANNELS = 16,
    ADD_CHANNELS= 8,
    DELETE_POSTS = 4,
    INVITE = 2,
    POST = 1,
    NONE = 0
}


export class ServerPermissions {

    permissions:number = 0;
    constructor(permissions: number) {

    }

    canInvite() :boolean{
        return (this.permissions & Permission.INVITE) == Permission.INVITE;
    }

    canPost() :boolean{
        return (this.permissions & Permission.POST) == Permission.POST;
    }
}