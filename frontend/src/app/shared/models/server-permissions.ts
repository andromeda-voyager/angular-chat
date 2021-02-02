enum Permission {
    FULL = 255,
    MANAGE_SERVER = 128,
    MANAGE_ROLES = 64,
    MANAGE_CHANNELS = 32,
    MANAGE_MEMBERS = 16,
    MANAGE_MESSAGES = 8,
    UNUSED = 4,
    INVITE = 2,
    POST = 1,
    NONE = 0
}


export class ServerPermissions {

    permissions: number = 0;
    constructor(permissions: number) {

    }

    canInvite(): boolean {
        return (this.permissions & Permission.INVITE) == Permission.INVITE;
    }

    canPost(): boolean {
        return (this.permissions & Permission.POST) == Permission.POST;
    }
}