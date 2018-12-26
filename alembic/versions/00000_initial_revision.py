"""initial revision

Revision ID: 51638789f439
Revises:
Create Date: 2018-12-21 17:07:26.432233

"""
from alembic import op, context
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision = '51638789f439'
down_revision = None
branch_labels = None
depends_on = None


def upgrade():
    op.create_table(
        "town_node",
        sa.Column('id', sa.Integer, primary_key=True),
        sa.Column('pub_key', sa.Text, nullable=False),
        sa.Column('amount_collected', sa.Integer, nullable=False),
        sa.Column('amount_received', sa.Integer, nullable=False),

        sa.Column('created_at', sa.DateTime(timezone=True), server_default=sa.func.now(), nullable=False),
        sa.Column('updated_at', sa.DateTime(timezone=True), nullable=True),
    )
    op.create_unique_constraint('uq_node_pub_key', 'town_node', ['pub_key'])

    op.create_table(
        "town_address",
        sa.Column('id', sa.Integer, primary_key=True),
        sa.Column('value', sa.String(35), nullable=False),
        sa.Column('amount_collected', sa.Integer, nullable=False),
        sa.Column('amount_received', sa.Integer, nullable=False),

        sa.Column('created_at', sa.DateTime(timezone=True), server_default=sa.func.now(), nullable=False),
        sa.Column('updated_at', sa.DateTime(timezone=True), nullable=True),
    )
    op.create_unique_constraint('uq_address_value', 'town_address', ['value'])

    op.create_table(
        "town_slug",
        sa.Column('id', sa.Integer, primary_key=True),
        sa.Column('slug', sa.String(100), nullable=False),

        sa.Column('current_id', sa.Integer, sa.ForeignKey('town_slug.id', name=op.f('town_slug_current_id_fkey')), nullable=True),
        sa.Column('created_at', sa.DateTime(timezone=True), server_default=sa.func.now(), nullable=False),
        sa.Column('updated_at', sa.DateTime(timezone=True), nullable=True),
    )
    op.create_index(op.f('town_article_current_id'), 'town_slug', ['current_id'])
    op.create_unique_constraint('uq_slug', 'town_slug', ['slug'])

    op.create_table(
        "town_article",
        sa.Column('id', sa.Integer, primary_key=True),
        sa.Column('title', sa.String(100), nullable=False),
        sa.Column('slug', sa.String(100), nullable=False),
        sa.Column('lang', sa.Integer, nullable=False),
        sa.Column('amount_collected', sa.Integer, nullable=False),
        sa.Column('amount_received', sa.Integer, nullable=False),
        sa.Column('status', sa.Integer, nullable=False),
        sa.Column('subtitle', sa.String(255), nullable=True),
        sa.Column('body_md', sa.Text, nullable=False),
        sa.Column('body_html', sa.Text, nullable=False),

        sa.Column('address_id', sa.Integer, sa.ForeignKey('town_address.id', name=op.f('town_address_id_fkey')), nullable=True),
        sa.Column('node_id', sa.Integer, sa.ForeignKey('town_node.id', name=op.f('town_node_id_fkey')), nullable=True),

        sa.Column('created_at', sa.DateTime(timezone=True), server_default=sa.func.now(), nullable=False),
        sa.Column('updated_at', sa.DateTime(timezone=True), nullable=True),
        sa.Column('published_at', sa.DateTime(timezone=True), nullable=True),
    )
    op.create_index(op.f('town_article_address_id'), 'town_article', ['address_id'])
    op.create_index(op.f('town_article_node_id'), 'town_article', ['node_id'])
    op.create_index(op.f('town_article_slug'), 'town_article', ['slug'])
    op.create_unique_constraint('uq_article_slug', 'town_article', ['slug'])

    if 'with-bootstrap' in context.get_x_argument(as_dictionary=False):
        from bootstrap.utils import insert_data
        from bootstrap.slugs import slugs
        from bootstrap.addresses import addresses
        from bootstrap.articles import articles

        insert_data("town_address", addresses)
        insert_data("town_slug", slugs)
        insert_data("town_article", articles)


def downgrade():
    op.drop_table('town_slug')
    op.drop_table('town_article')
