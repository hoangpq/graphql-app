import React from 'react';

class Migration extends React.Component {
  constructor(props) {
    super(props);
    this.actionStartMigration = this.actionStartMigration.bind(this);
  }

  actionStartMigration() {
    // TODO
  }

  render() {
    return (
      <div>
        <button onClick={this.actionStartMigration}>Start migration</button>
      </div>
    )
  }
}

export default Migration;
