from backend.local import LocalBackend
from backends import provider
from channel import Channel
from registry.local_client import LocalRegistryClient


class ChannelManager(object):
    def __init__(self, backend, job, role, channels_roles):
        self.job = job
        self.role = role
        self.channels_roles = channels_roles

        self.channels = {}

        self.registry_client = LocalRegistryClient()
        self.backend = provider.get(backend)

    def join(self, name):
        '''
        joins a channel
        '''
        if self.join_done(name):
            return True

        ret = await self.registry_client.register(
            self.job, name, self.role, self.backend.uid(), self.backend.endpoint()
        )

        if ret is True:
            self.channels[name] = Channel(name, self.backend)

        # role_map should have at most two entries
        role_map = self.channels_roles[name]

        channel_info = self.registry_client.get(self.job, name)
        for role, end_id, endpoint in channel_info:
            # the same backend id; skip
            if end_id is self.backend.id():
                continue

            proceed = False
            for from_role, to_role in role_map.items():
                proceed = self.role is from_role and role is to_role

            if not proceed:
                continue

            # connect to endpoint
            self.backend.connect(end_id, endpoint)

            # notify end_id of the channel handled by the backend
            self.backend.notify(end_id, name)

            # update channel
            self.channels[name].add(end_id)

        return ret

    def leave(self, name):
        '''
        leave a channel
        '''
        if not self.join_done(name):
            return

        ret = await self.registry_client.reset_channel(
            self.job, name, self.role, self.backend.uid()
        )

        del self.channels[name]

        return ret

    def join_done(self, name):
        '''
        check if node joined a channel or not
        '''
        if name in self.channels:
            return True
        else:
            return False

    def get(self, name):
        '''
        returns a channel object in the given channel
        '''
        if not self.join_done(name):
            # didn't join the channel yet
            return None

        return self.channels[name]

    def _update(self, name, end_id):
        '''
        add an end ID to a channel with 'name'
        '''
        if not self.join_done(name):
            # didn't join the channel yet
            return

        channel = self.channels[name]

        channel.add(end_id)
